package cloudrun

import (
	"context"
	"encoding/json"
	"github.com/cloudquery/faker/v3"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	"google.golang.org/api/run/v1"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func createServicesServer() (*client.Services, error) {
	ctx := context.Background()
	var service run.Service
	if err := faker.FakeData(&service); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v1/projects/testProject/locations/-/services", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &run.ListServicesResponse{
			Items: []*run.Service{
				&service,
			},
			Metadata: nil,
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
	svc, err := run.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		CloudRun: svc,
	}, nil
}

func TestServices(t *testing.T) {
	client.GcpMockTestHelper(t, Services(), createServicesServer, client.TestOptions{})
}
