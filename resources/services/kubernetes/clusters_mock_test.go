package kubernetes

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cloudquery/cq-provider-gcp/client"
	faker "github.com/cloudquery/faker/v3"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"

	"google.golang.org/api/container/v1"
)

func createClusters() (*client.Services, error) {
	ctx := context.Background()
	var cluster container.Cluster
	if err := faker.FakeData(&cluster); err != nil {
		return nil, err
	}
	mux := httprouter.New()

	cluster.CreateTime = time.Now().Format(time.RFC3339)
	cluster.ExpireTime = time.Now().Format(time.RFC3339)
	cluster.Endpoint = "192.168.0.1"

	cluster.MaintenancePolicy.Window.RecurringWindow.Window.StartTime = time.Now().Format(time.RFC3339)
	cluster.MaintenancePolicy.Window.RecurringWindow.Window.EndTime = time.Now().Format(time.RFC3339)

	cluster.NodePools[0].Management.UpgradeOptions.AutoUpgradeStartTime = time.Now().Format(time.RFC3339)

	mux.GET("/v1/projects/testProject/locations/-/clusters", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &container.ListClustersResponse{
			Clusters: []*container.Cluster{&cluster},
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
	svc, err := container.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Container: svc,
	}, nil
}

func TestClusters(t *testing.T) {
	client.GcpMockTestHelper(t, Clusters(), createClusters, client.TestOptions{})
}
