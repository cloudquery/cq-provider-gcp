package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	logging2 "google.golang.org/api/logging/v2"
	"google.golang.org/api/option"
)

func createLoggingSinks() (*logging2.Service, error) {
	ctx := context.Background()
	var logSink logging2.LogSink
	if err := faker.FakeData(&logSink); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v2/projects/testProject/sinks", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &logging2.ListSinksResponse{
			Sinks: []*logging2.LogSink{&logSink},
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
	svc, err := logging2.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func TestLoggingSinks(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: LoggingSinks(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			loggingSvc, err := createLoggingSinks()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Logging: loggingSvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, Provider, resource)
}
