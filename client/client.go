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

	if len(providerConfig.Folders) > 0 {
		logger.Debug("Adding projects in specified folders", "folders", providerConfig.Folders)
		proj, err := getProjects(logger, providerConfig.Folders, options...)
		if err != nil {
			return nil, err
		}
		appendWithoutDupes(&projects, proj)
	}
	if providerConfig.FolderQuery != "" {
		logger.Debug("Querying folders", "query", providerConfig.FolderQuery)
		queriedFolders, err := getFolders(logger, providerConfig.FolderQuery, options...)
		if err != nil {
			return nil, err
		}
		proj, err := getProjects(logger, queriedFolders, options...)
		if err != nil {
			return nil, err
		}
		appendWithoutDupes(&projects, proj)
	}
	if len(projects) == 0 {
		logger.Info("No project_ids specified, assuming all active projects")
		var err error
		projects, err = getProjects(logger, []string{""}, options...)
		if err != nil {
			return nil, err
		}
	}

	if err := validateProjects(projects); err != nil {
		return nil, err
	}
	services, err := initServices(context.Background(), options)
	if err != nil {
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

func getProjects(logger hclog.Logger, folders []string, options ...option.ClientOption) ([]string, error) {
	if len(folders) == 0 {
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	service, err := cloudresourcemanager.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}

	var (
		projects []string
		inactive int
	)

	for _, folder := range folders {
		call := service.Projects.List()
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

func getFolders(logger hclog.Logger, query string, options ...option.ClientOption) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	service, err := cloudresourcemanager.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}

	var folders []string

	call := service.Folders.Search().Query(query)
	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, folder := range output.Folders {
			if folder.State == "ACTIVE" {
				folders = append(folders, folder.Name)
			} else {
				logger.Info("Folder state is not active. Folder will be ignored", "folder_name", folder.Name, "folder_state", folder.State)
			}
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	logger.Debug("Search query found folders", "folders", folders)

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
