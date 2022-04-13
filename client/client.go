package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
)

const defaultProjectIdName = "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"

const serviceAccountEnvKey = "CQ_SERVICE_ACCOUNT_KEY_JSON"

type Client struct {
	projects []string
	logger   hclog.Logger
	backoff  BackoffSettings

	// All gcp services initialized by client
	Services *Services
	// this is set by table client multiplexer
	ProjectId string
}

func NewGcpClient(log hclog.Logger, bo BackoffSettings, projects []string, services *Services) *Client {
	return &Client{
		projects: projects,
		logger:   log,
		backoff:  bo,
		Services: services,
	}
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

// withProject allows multiplexer to create a new client with given subscriptionId
func (c Client) withProject(project string) *Client {
	return &Client{
		backoff:   c.backoff,
		projects:  c.projects,
		Services:  c.Services,
		logger:    c.logger.With("project_id", project),
		ProjectId: project,
	}
}

func isValidJson(content []byte) error {
	var v map[string]interface{}
	err := json.Unmarshal(content, &v)
	if err != nil {
		var syntaxError *json.SyntaxError
		if errors.As(err, &syntaxError) {
			return fmt.Errorf("the environment variable %s should contain valid JSON object. %w", serviceAccountEnvKey, err)
		}
		return err
	}
	return nil
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	projects := providerConfig.ProjectIDs
	if providerConfig.FolderMaxDepth == 0 {
		providerConfig.FolderMaxDepth = 5
	}

	serviceAccountKeyJSON := []byte(providerConfig.ServiceAccountKeyJSON)
	if len(serviceAccountKeyJSON) == 0 {
		serviceAccountKeyJSON = []byte(os.Getenv(serviceAccountEnvKey))
	}

	// Add a fake request reason because it is not possible to pass nil options
	options := append([]option.ClientOption{option.WithRequestReason("cloudquery resource fetch")}, providerConfig.ClientOptions()...)
	if len(serviceAccountKeyJSON) != 0 {
		if err := isValidJson(serviceAccountKeyJSON); err != nil {
			return nil, err
		}
		options = append(options, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	if providerConfig.ProjectFilter != "" {
		return nil, errors.New("ProjectFilter config option is deprecated")
	}

	services, err := initServices(context.Background(), options)
	if err != nil {
		return nil, err
	}

	if len(providerConfig.Folders) > 0 {
		logger.Debug("Listing folders", "folders", providerConfig.Folders)

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		var folderList []string
		for _, f := range folderList {
			folderAndChildren, err := listFolders(ctx, logger, services.ResourceManager, f, int(providerConfig.FolderMaxDepth))
			if err != nil {
				return nil, err
			}
			folderList = append(folderList, folderAndChildren...)
		}
		proj, err := getProjects(logger, services.ResourceManager, folderList)
		if err != nil {
			return nil, err
		}
		appendWithoutDupes(&projects, proj)
	}
	if len(projects) == 0 {
		logger.Info("No project_ids specified, assuming all active projects")
		var err error
		projects, err = getProjects(logger, services.ResourceManager, nil)
		if err != nil {
			return nil, err
		}
	}

	if err := validateProjects(projects); err != nil {
		return nil, err
	}

	client := NewGcpClient(logger, providerConfig.Backoff(), projects, services)
	return client, nil
}

func validateProjects(projects []string) error {
	for _, project := range projects {
		if project == defaultProjectIdName {
			return fmt.Errorf("please specify a valid project_id in config.hcl instead of <CHANGE_THIS_TO_YOUR_PROJECT_ID>")
		}
	}
	return nil
}

func getProjects(logger hclog.Logger, service *cloudresourcemanager.Service, folders []string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if len(folders) == 0 {
		folders = []string{""}
	}

	var (
		projects []string
		inactive int
	)

	for _, folder := range folders {
		call := service.Projects.List().Context(ctx)
		if folder != "" {
			call.Parent(folder)
		}

		for {
			output, err := call.Do()
			if err != nil {
				return nil, err
			}
			for _, project := range output.Projects {
				if project.State == "ACTIVE" {
					projects = append(projects, project.ProjectId)
				} else {
					logger.Info("Project state is not active. Project will be ignored", "project_id", project.ProjectId, "project_state", project.State)
					inactive++
				}
			}
			if output.NextPageToken == "" {
				break
			}
			call.PageToken(output.NextPageToken)
		}
	}

	if len(projects) == 0 {
		if inactive > 0 {
			return nil, fmt.Errorf("project listing failed: no active projects")
		}
		return nil, fmt.Errorf("project listing failed")
	}

	return projects, nil
}

func listFolders(ctx context.Context, logger hclog.Logger, service *cloudresourcemanager.Service, parent string, maxDepth int) ([]string, error) {
	folders := []string{
		parent,
	}
	if maxDepth < 0 {
		return folders, nil
	}

	call := service.Folders.List().Context(ctx).Parent("folders/" + parent)
	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, folder := range output.Folders {
			if folder.State != "ACTIVE" {
				logger.Info("Folder state is not active. Folder will be ignored", "folder_name", folder.Name, "folder_state", folder.State)
				continue
			}
			fList, err := listFolders(ctx, logger, service, folder.Name, maxDepth-1)
			if err != nil {
				return nil, err
			}
			folders = append(folders, fList...)
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	logger.Debug("List query found folders", "folders", folders)

	return folders, nil
}

func appendWithoutDupes(dst *[]string, src []string) {
	dstMap := make(map[string]struct{}, len(*dst))
	for i := range *dst {
		dstMap[(*dst)[i]] = struct{}{}
	}
	for i := range src {
		if _, ok := dstMap[src[i]]; ok {
			continue
		}
		dstMap[src[i]] = struct{}{}
		*dst = append(*dst, src[i])
	}
}
