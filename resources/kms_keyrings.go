package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/cloudkms/v1"
)

func KmsKeyrings() *schema.Table {
	return &schema.Table{
		Name:                 "gcp_kms_keyrings",
		Resolver:             fetchKmsKeyrings,
		Multiplex:            client.ProjectMultiplex,
		IgnoreError:          client.IgnoreErrorHandler,
		DeleteFilter:         client.DeleteProjectFilter,
		PostResourceResolver: client.AddGcpMetadata,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: resolveKmsKeyringCreateTime,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:                 "gcp_kms_keyring_crypto_keys",
				Resolver:             fetchKmsKeyringCryptoKeys,
				IgnoreError:          client.IgnoreErrorHandler,
				PostResourceResolver: client.AddGcpMetadata,
				Columns: []schema.Column{
					{
						Name:     "keyring_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "project_id",
						Type:     schema.TypeString,
						Resolver: client.ResolveProject,
					},
					{
						Name: "location",
						Type: schema.TypeString,
					},
					{
						Name:     "create_time",
						Type:     schema.TypeTimestamp,
						Resolver: resolveKmsKeyringCryptoKeyCreateTime,
					},
					{
						Name: "labels",
						Type: schema.TypeJSON,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name:     "next_rotation_time",
						Type:     schema.TypeTimestamp,
						Resolver: resolveKmsKeyringCryptoKeyNextRotationTime,
					},
					{
						Name:     "primary_algorithm",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.Algorithm"),
					},
					{
						Name:     "primary_attestation_cert_chains_cavium_certs",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Primary.Attestation.CertChains.CaviumCerts"),
					},
					{
						Name:     "primary_attestation_cert_chains_google_card_certs",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Primary.Attestation.CertChains.GoogleCardCerts"),
					},
					{
						Name:     "primary_attestation_cert_chains_google_partition_certs",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Primary.Attestation.CertChains.GooglePartitionCerts"),
					},
					{
						Name:     "primary_attestation_content",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.Attestation.Content"),
					},
					{
						Name:     "primary_attestation_format",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.Attestation.Format"),
					},
					{
						Name:     "primary_create_time",
						Type:     schema.TypeTimestamp,
						Resolver: resolveKmsKeyringCryptoKeyPrimaryCreateTime,
					},
					{
						Name:     "primary_destroy_event_time",
						Type:     schema.TypeTimestamp,
						Resolver: resolveKmsKeyringCryptoKeyPrimaryDestroyEventTime,
					},
					{
						Name:     "primary_destroy_time",
						Type:     schema.TypeTimestamp,
						Resolver: resolveKmsKeyringCryptoKeyPrimaryDestroyTime,
					},
					{
						Name:     "primary_external_protection_level_options_external_key_uri",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.ExternalProtectionLevelOptions.ExternalKeyUri"),
					},
					{
						Name:     "primary_generate_time",
						Type:     schema.TypeTimestamp,
						Resolver: resolveKmsKeyringCryptoKeyPrimaryGenerateTime,
					},
					{
						Name:     "primary_import_failure_reason",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.ImportFailureReason"),
					},
					{
						Name:     "primary_import_job",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.ImportJob"),
					},
					{
						Name:     "primary_import_time",
						Type:     schema.TypeTimestamp,
						Resolver: resolveKmsKeyringCryptoKeyPrimaryImportTime,
					},
					{
						Name:     "primary_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.Name"),
					},
					{
						Name:     "primary_protection_level",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.ProtectionLevel"),
					},
					{
						Name:     "primary_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Primary.State"),
					},
					{
						Name: "purpose",
						Type: schema.TypeString,
					},
					{
						Name: "rotation_period",
						Type: schema.TypeString,
					},
					{
						Name:     "version_template_algorithm",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VersionTemplate.Algorithm"),
					},
					{
						Name:     "version_template_protection_level",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VersionTemplate.ProtectionLevel"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchKmsKeyrings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	locations, err := getAllKmsLocations(ctx, c.ProjectId, c.Services.Kms)
	if err != nil {
		return fmt.Errorf("failed to get kms locations. %w", err)
	}
	nextPageToken := ""
	for _, l := range locations {
		call := c.Services.Kms.Projects.Locations.KeyRings.List(l.Name).Context(ctx)
		for {
			call.PageToken(nextPageToken)
			resp, err := call.Do()
			if err != nil {
				return err
			}
			res <- resp.KeyRings

			if resp.NextPageToken == "" {
				break
			}
			nextPageToken = resp.NextPageToken
		}
	}
	return nil
}
func resolveKmsKeyringCreateTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudkms.KeyRing)
	if !ok {
		return fmt.Errorf("expected *cloudkms.KeyRing but got %T", p)
	}
	date, err := client.ParseISODate(p.CreateTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func fetchKmsKeyringCryptoKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	keyRing, ok := parent.Item.(*cloudkms.KeyRing)
	if !ok {
		return fmt.Errorf("expected *cloudkms.KeyRing but got %T", keyRing)
	}
	nextPageToken := ""
	call := c.Services.Kms.Projects.Locations.KeyRings.CryptoKeys.List(keyRing.Name).Context(ctx)
	for {
		call.PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return err
		}
		res <- resp.CryptoKeys

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return nil
}
func resolveKmsKeyringCryptoKeyCreateTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudkms.CryptoKey)
	if !ok {
		return fmt.Errorf("expected *cloudkms.CryptoKey but got %T", p)
	}
	date, err := client.ParseISODate(p.CreateTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveKmsKeyringCryptoKeyNextRotationTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudkms.CryptoKey)
	if !ok {
		return fmt.Errorf("expected *cloudkms.CryptoKey but got %T", p)
	}
	date, err := client.ParseISODate(p.NextRotationTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveKmsKeyringCryptoKeyPrimaryCreateTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudkms.CryptoKey)
	if !ok {
		return fmt.Errorf("expected *cloudkms.CryptoKey but got %T", p)
	}
	if p.Primary == nil {
		return nil
	}
	date, err := client.ParseISODate(p.Primary.CreateTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveKmsKeyringCryptoKeyPrimaryDestroyEventTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudkms.CryptoKey)
	if !ok {
		return fmt.Errorf("expected *cloudkms.CryptoKey but got %T", p)
	}
	if p.Primary == nil {
		return nil
	}
	date, err := client.ParseISODate(p.Primary.DestroyEventTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveKmsKeyringCryptoKeyPrimaryDestroyTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudkms.CryptoKey)
	if !ok {
		return fmt.Errorf("expected *cloudkms.CryptoKey but got %T", p)
	}
	if p.Primary == nil {
		return nil
	}
	date, err := client.ParseISODate(p.Primary.DestroyTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveKmsKeyringCryptoKeyPrimaryGenerateTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudkms.CryptoKey)
	if !ok {
		return fmt.Errorf("expected *cloudkms.CryptoKey but got %T", p)
	}
	if p.Primary == nil {
		return nil
	}
	date, err := client.ParseISODate(p.Primary.GenerateTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}
func resolveKmsKeyringCryptoKeyPrimaryImportTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*cloudkms.CryptoKey)
	if !ok {
		return fmt.Errorf("expected *cloudkms.CryptoKey but got %T", p)
	}
	if p.Primary == nil {
		return nil
	}
	date, err := client.ParseISODate(p.Primary.ImportTime)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, date)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func getAllKmsLocations(ctx context.Context, projectId string, kms *cloudkms.Service) ([]*cloudkms.Location, error) {
	var locations []*cloudkms.Location
	call := kms.Projects.Locations.List("projects/" + projectId).Context(ctx)
	nextPageToken := ""
	for {
		call.PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return nil, err
		}
		locations = append(locations, resp.Locations...)

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return locations, nil
}
