package client

import (
	"context"

	"cloud.google.com/go/bigquery"
	compute "cloud.google.com/go/compute/apiv1"
	domains "cloud.google.com/go/domains/apiv1beta1"
	functions "cloud.google.com/go/functions/apiv1"
	iam "cloud.google.com/go/iam/admin/apiv1"
	"cloud.google.com/go/logging/logadmin"
	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/storage"
	kms "google.golang.org/api/cloudkms/v1"
	dns "google.golang.org/api/dns/v1"
	sql "google.golang.org/api/sqladmin/v1beta4"

	"google.golang.org/api/option"
)

type Services struct {
	Kms            *kms.Service
	Storage        *storage.Client
	Sql            *sql.Service
	Iam            *iam.IamClient
	CloudFunctions *functions.CloudFunctionsClient
	Domain         *domains.Client
	Compute        *compute.ProjectsClient
	BigQuery       bigqueryInstantiator
	Dns            *dns.Service
	Logging        loggingInstantiator
	Monitoring     *monitoring.ServiceMonitoringClient

	ResourceManagerFolders  *resourcemanager.FoldersClient
	ResourceManagerProjects *resourcemanager.ProjectsClient
}

type (
	bigqueryInstantiator func(projectID string) (*bigquery.Client, error)
	loggingInstantiator  func(parent string) (*logadmin.Client, error)
)

func initServices(ctx context.Context, options []option.ClientOption) (*Services, error) {
	kmsSvc, err := kms.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	storageSvc, err := storage.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	sqlSvc, err := sql.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	iamSvc, err := iam.NewIamClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	cfSvc, err := functions.NewCloudFunctionsClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	domainSvc, err := domains.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	computeSvc, err := compute.NewProjectsRESTClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	dnsSvc, err := dns.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	monitoringSvc, err := monitoring.NewServiceMonitoringClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	resourceManagerFolders, err := resourcemanager.NewFoldersClient(ctx, options...)
	if err != nil {
		return nil, err
	}
	resourceManagerProjects, err := resourcemanager.NewProjectsClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &Services{
		Kms:                     kmsSvc,
		Storage:                 storageSvc,
		Sql:                     sqlSvc,
		Iam:                     iamSvc,
		CloudFunctions:          cfSvc,
		Domain:                  domainSvc,
		Compute:                 computeSvc,
		BigQuery:                bigqueryClient(ctx, options...),
		Dns:                     dnsSvc,
		Logging:                 loggingClient(ctx, options...),
		Monitoring:              monitoringSvc,
		ResourceManagerFolders:  resourceManagerFolders,
		ResourceManagerProjects: resourceManagerProjects,
	}, nil
}

func bigqueryClient(ctx context.Context, options ...option.ClientOption) bigqueryInstantiator {
	return func(projectID string) (*bigquery.Client, error) {
		return bigquery.NewClient(ctx, projectID, options...)
	}
}

func loggingClient(ctx context.Context, options ...option.ClientOption) loggingInstantiator {
	return func(parent string) (*logadmin.Client, error) {
		return logadmin.NewClient(ctx, parent, options...)
	}
}
