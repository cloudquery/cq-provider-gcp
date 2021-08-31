package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/bigquery/v2"
	"google.golang.org/api/option"
)

func TestBigqueryDatasets(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: BigqueryDatasets(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			bigquerySvc, err := createBigqueryDatasets()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				BigQuery: bigquerySvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, Provider, resource)
}

func createBigqueryDatasets() (*bigquery.Service, error) {
	id := "testDataset"
	mux := httprouter.New()
	var dataset bigquery.Dataset
	if err := faker.FakeData(&dataset); err != nil {
		return nil, err
	}
	dataset.Id = id
	dataset.DatasetReference = &bigquery.DatasetReference{
		DatasetId: id,
	}
	mux.GET("/projects/testProject/datasets", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &bigquery.DatasetList{
			Datasets: []*bigquery.DatasetListDatasets{
				{
					DatasetReference: dataset.DatasetReference,
				},
			},
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

	mux.GET("/projects/testProject/datasets/testDataset", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&dataset)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	mux.GET("/projects/testProject/datasets/testDataset/tables", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &bigquery.TableList{

			Tables: []*bigquery.TableListTables{
				{
					Id: id,
					TableReference: &bigquery.TableReference{
						TableId: id,
					},
				},
			},
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

	var table bigquery.Table
	table.Id = id
	table.TableReference = &bigquery.TableReference{
		TableId: id,
	}
	if err := faker.FakeData(&table.Model); err != nil {
		return nil, err
	}
	if err := faker.FakeData(&table.View); err != nil {
		return nil, err
	}
	if err := faker.FakeData(&table.Type); err != nil {
		return nil, err
	}
	schema := bigquery.TableSchema{
		Fields: []*bigquery.TableFieldSchema{{
			Name: "test",
			Type: "test",
		},
		},
	}
	table.Schema = &schema

	table.ExternalDataConfiguration = &bigquery.ExternalDataConfiguration{
		Autodetect: true,
		Schema:     &schema,
		SourceUris: []string{"test"},
	}
	table.Labels = map[string]string{
		"test": "test",
	}
	table.Clustering = &bigquery.Clustering{
		Fields: []string{"test"},
	}
	if err := faker.FakeData(&table.Description); err != nil {
		return nil, err
	}
	if err := faker.FakeData(&table.EncryptionConfiguration); err != nil {
		return nil, err
	}

	mux.GET("/projects/testProject/datasets/testDataset/tables/testDataset", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&table)
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
	svc, err := bigquery.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc, nil
}
