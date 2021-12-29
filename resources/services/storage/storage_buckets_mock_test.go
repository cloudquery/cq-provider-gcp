//go:build mock
// +build mock

package storage

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	faker "github.com/cloudquery/faker/v3"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	storage2 "google.golang.org/api/storage/v1"
)

func createStorageTestServer() (*client.Services, error) {
	ctx := context.Background()
	var bucket storage2.Bucket
	if err := faker.FakeData(&bucket); err != nil {
		return nil, err
	}
	bucket.Name = "testBucket"
	mux := httprouter.New()
	mux.GET("/b", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &storage2.Buckets{Items: []*storage2.Bucket{&bucket}}
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

	var policy storage2.Policy
	if err := faker.FakeData(&policy); err != nil {
		return nil, err
	}
	mux.GET("/b/testBucket/iam", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &policy
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
	svc, err := storage2.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Storage: svc,
	}, nil
}

func TestStorageBucket(t *testing.T) {
	client.GcpMockTestHelper(t, StorageBuckets(), createStorageTestServer, client.TestOptions{})
}
