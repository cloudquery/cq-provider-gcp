package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/cloudkms/v1"
)

func KmsKeyrings() *schema.Table {
	return &schema.Table{
		Name:                 "gcp_kms_keyrings",
		Description:          "KeyRing: A KeyRing is a toplevel logical grouping of CryptoKeys. ",
		Resolver:             fetchKmsKeyrings,
		Multiplex:            client.ProjectMultiplex,
		IgnoreError:          client.IgnoreErrorHandler,
		DeleteFilter:         client.DeleteProjectFilter,
		PostResourceResolver: client.AddGcpMetadata,
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name:        "create_time",
				Description: "CreateTime: Output only",
				Type:        schema.TypeTimestamp,
				Resolver:    client.ISODateResolver("CreateTime"),
			},
			{
				Name:        "name",
				Description: "Name: Output only",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:                 "gcp_kms_keyring_crypto_keys",
				Description:          "CryptoKey: A CryptoKey represents a logical key that can be used for cryptographic operations",
				Resolver:             fetchKmsKeyringCryptoKeys,
				IgnoreError:          client.IgnoreErrorHandler,
				PostResourceResolver: client.AddGcpMetadata,
				Columns: []schema.Column{
					{
						Name:        "keyring_id",
						Description: "Unique ID of gcp_kms_keyrings table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "project_id",
						Description: "GCP Project Id of the resource",
						Type:        schema.TypeString,
						Resolver:    client.ResolveProject,
					},
					{
						Name: "location",
						Type: schema.TypeString,
					},
					{
						Name:     "policy",
						Type:     schema.TypeJSON,
						Resolver: resolveKmsKeyringCryptoKeyPolicy,
					},
					{
						Name:        "create_time",
						Description: "CreateTime: Output only",
						Type:        schema.TypeTimestamp,
						Resolver:    client.ISODateResolver("CreateTime"),
					},
					{
						Name:        "labels",
						Description: "Labels: Labels with user-defined metadata",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "name",
						Description: "Name: Output only",
						Type:        schema.TypeString,
					},
					{
						Name:        "next_rotation_time",
						Description: "NextRotationTime: At next_rotation_time, the Key Management Service will automatically: 1",
						Type:        schema.TypeTimestamp,
						Resolver:    client.ISODateResolver("NextRotationTime"),
					},
					{
						Name:        "primary_algorithm",
						Description: "Algorithm: Output only",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Algorithm"),
					},
					{
						Name:        "primary_attestation_cert_chains_cavium_certs",
						Description: "CaviumCerts: Cavium certificate chain corresponding to the attestation.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Primary.Attestation.CertChains.CaviumCerts"),
					},
					{
						Name:        "primary_attestation_cert_chains_google_card_certs",
						Description: "GoogleCardCerts: Google card certificate chain corresponding to the attestation.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Primary.Attestation.CertChains.GoogleCardCerts"),
					},
					{
						Name:        "primary_attestation_cert_chains_google_partition_certs",
						Description: "GooglePartitionCerts: Google partition certificate chain corresponding to the attestation.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Primary.Attestation.CertChains.GooglePartitionCerts"),
					},
					{
						Name:        "primary_attestation_content",
						Description: "Content: Output only",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Attestation.Content"),
					},
					{
						Name:        "primary_attestation_format",
						Description: "Format: Output only",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Attestation.Format"),
					},
					{
						Name:        "primary_create_time",
						Description: "CreateTime: Output only",
						Type:        schema.TypeTimestamp,
						Resolver:    client.ISODateResolver("Primary.CreateTime"),
					},
					{
						Name:        "primary_destroy_event_time",
						Description: "DestroyEventTime: Output only",
						Type:        schema.TypeTimestamp,
						Resolver:    client.ISODateResolver("Primary.DestroyEventTime"),
					},
					{
						Name:        "primary_destroy_time",
						Description: "DestroyTime: Output only",
						Type:        schema.TypeTimestamp,
						Resolver:    client.ISODateResolver("Primary.DestroyTime"),
					},
					{
						Name:        "primary_external_protection_level_options_external_key_uri",
						Description: "ExternalKeyUri: The URI for an external resource that this CryptoKeyVersion represents.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.ExternalProtectionLevelOptions.ExternalKeyUri"),
					},
					{
						Name:        "primary_generate_time",
						Description: "GenerateTime: Output only",
						Type:        schema.TypeTimestamp,
						Resolver:    client.ISODateResolver("Primary.GenerateTime"),
					},
					{
						Name:        "primary_import_failure_reason",
						Description: "ImportFailureReason: Output only",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.ImportFailureReason"),
					},
					{
						Name:        "primary_import_job",
						Description: "ImportJob: Output only",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.ImportJob"),
					},
					{
						Name:        "primary_import_time",
						Description: "ImportTime: Output only",
						Type:        schema.TypeTimestamp,
						Resolver:    client.ISODateResolver("Primary.ImportTime"),
					},
					{
						Name:        "primary_name",
						Description: "Name: Output only",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.Name"),
					},
					{
						Name:        "primary_protection_level",
						Description: "ProtectionLevel: Output only",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.ProtectionLevel"),
					},
					{
						Name:        "primary_state",
						Description: "State: The current state of the CryptoKeyVersion.  Possible values:   \"CRYPTO_KEY_VERSION_STATE_UNSPECIFIED\" - Not specified.   \"PENDING_GENERATION\" - This version is still being generated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Primary.State"),
					},
					{
						Name:        "purpose",
						Description: "Purpose: Immutable",
						Type:        schema.TypeString,
					},
					{
						Name:        "rotation_period",
						Description: "RotationPeriod: next_rotation_time will be advanced by this period when the service automatically rotates a key",
						Type:        schema.TypeString,
					},
					{
						Name:        "version_template_algorithm",
						Description: "Algorithm: Required",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionTemplate.Algorithm"),
					},
					{
						Name:        "version_template_protection_level",
						Description: "ProtectionLevel: ProtectionLevel to use when creating a CryptoKeyVersion based on this template",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VersionTemplate.ProtectionLevel"),
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
func resolveKmsKeyringCryptoKeyPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	client := meta.(*client.Client)
	p, ok := resource.Item.(*cloudkms.CryptoKey)
	if !ok {
		return fmt.Errorf("expected *cloudkms.CryptoKey but got %T", p)
	}
	call := client.Services.Kms.Projects.Locations.KeyRings.CryptoKeys.
		GetIamPolicy(p.Name).
		Context(ctx)
	resp, err := call.Do()
	if err != nil {
		return err
	}

	var policy map[string]interface{}
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &policy); err != nil {
		return err
	}

	return resource.Set(c.Name, policy)
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
