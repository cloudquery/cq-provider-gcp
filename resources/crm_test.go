package resources_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-gcp/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
)

func TestCrmProjects(t *testing.T) {
	resource := ResourceTestData{
		Table: resources.CrmProjects(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
			Resources:  []client.Resource{{Name: "crm.projects"}},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			crmSvc, err := createCrmTestServer()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Crm: crmSvc,
			})
			return c, nil
		},
	}
	testResource(t, resources.Provider, resource)
}

func createCrmTestServer() (*cloudresourcemanager.Service, error) {
	ctx := context.Background()
	var project cloudresourcemanager.Project
	if err := faker.FakeData(&project); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v3/projects", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &cloudresourcemanager.ListProjectsResponse{
			Projects: []*cloudresourcemanager.Project{&project},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	svc, err := cloudresourcemanager.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc, nil
}
