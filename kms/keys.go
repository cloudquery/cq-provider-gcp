package kms

import (
	"go.uber.org/zap"
	"google.golang.org/api/cloudkms/v1"
)

type CryptoKey struct {
	ID               uint `gorm:"primarykey"`
	ProjectID        string
	CreateTime       string
	Name             string
	NextRotationTime string

	PrimaryAlgorithm string

	CaviumCerts          []*CryptoKeyCertificateChainsCaviumCerts          `gorm:"constraint:OnDelete:CASCADE;foreignKey:ID"`
	GoogleCardCerts      []*CryptoKeyCertificateChainsGoogleCardCerts      `gorm:"constraint:OnDelete:CASCADE;foreignKey:ID"`
	GooglePartitionCerts []*CryptoKeyCertificateChainsGooglePartitionCerts `gorm:"constraint:OnDelete:CASCADE;foreignKey:ID"`

	PrimaryAttestationContent string
	PrimaryAttestationFormat  string

	PrimaryCreateTime       string
	PrimaryDestroyEventTime string
	PrimaryDestroyTime      string

	PrimaryExternalProtectionLevelOptionsExternalKeyUri string

	PrimaryGenerateTime        string
	PrimaryImportFailureReason string
	PrimaryImportJob           string
	PrimaryImportTime          string
	PrimaryName                string
	PrimaryProtectionLevel     string
	PrimaryState               string

	Purpose        string
	RotationPeriod string

	VersionTemplateAlgorithm       string
	VersionTemplateProtectionLevel string
}

func (CryptoKey) TableName() string {
	return "gcp_kms_crypto_keys"
}

type CryptoKeyCertificateChainsCaviumCerts struct {
	ID                           uint `gorm:"primarykey"`
	CryptoKeyCertificateChainsID uint
	Value                        string
}

func (CryptoKeyCertificateChainsCaviumCerts) TableName() string {
	return "gcp_kms_crypto_key_certificate_chains_cavium_certs"
}

type CryptoKeyCertificateChainsGoogleCardCerts struct {
	ID                           uint `gorm:"primarykey"`
	CryptoKeyCertificateChainsID uint
	Value                        string
}

func (CryptoKeyCertificateChainsGoogleCardCerts) TableName() string {
	return "gcp_kms_crypto_key_certificate_chains_google_card_certs"
}

type CryptoKeyCertificateChainsGooglePartitionCerts struct {
	ID                           uint `gorm:"primarykey"`
	CryptoKeyCertificateChainsID uint
	Value                        string
}

func (CryptoKeyCertificateChainsGooglePartitionCerts) TableName() string {
	return "gcp_kms_crypto_key_certificate_chains_google_partition_certs"
}

func (c *Client) transformCryptoKeys(values []*cloudkms.CryptoKey) []*CryptoKey {
	var tValues []*CryptoKey
	for _, value := range values {
		tValue := CryptoKey{
			ProjectID:        c.projectID,
			CreateTime:       value.CreateTime,
			Name:             value.Name,
			NextRotationTime: value.NextRotationTime,
			Purpose:          value.Purpose,
			RotationPeriod:   value.RotationPeriod,
		}
		if value.Primary != nil {

			tValue.PrimaryAlgorithm = value.Primary.Algorithm
			tValue.PrimaryCreateTime = value.Primary.CreateTime
			tValue.PrimaryDestroyEventTime = value.Primary.DestroyEventTime
			tValue.PrimaryDestroyTime = value.Primary.DestroyTime
			tValue.PrimaryGenerateTime = value.Primary.GenerateTime
			tValue.PrimaryImportFailureReason = value.Primary.ImportFailureReason
			tValue.PrimaryImportJob = value.Primary.ImportJob
			tValue.PrimaryImportTime = value.Primary.ImportTime
			tValue.PrimaryName = value.Primary.Name
			tValue.PrimaryProtectionLevel = value.Primary.ProtectionLevel
			tValue.PrimaryState = value.Primary.State

		}
		if value.VersionTemplate != nil {

			tValue.VersionTemplateAlgorithm = value.VersionTemplate.Algorithm
			tValue.VersionTemplateProtectionLevel = value.VersionTemplate.ProtectionLevel

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformCryptoKeyCertificateChainsCaviumCerts(values []string) []*CryptoKeyCertificateChainsCaviumCerts {
	var tValues []*CryptoKeyCertificateChainsCaviumCerts
	for _, v := range values {
		tValues = append(tValues, &CryptoKeyCertificateChainsCaviumCerts{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformCryptoKeyCertificateChainsGoogleCardCerts(values []string) []*CryptoKeyCertificateChainsGoogleCardCerts {
	var tValues []*CryptoKeyCertificateChainsGoogleCardCerts
	for _, v := range values {
		tValues = append(tValues, &CryptoKeyCertificateChainsGoogleCardCerts{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformCryptoKeyCertificateChainsGooglePartitionCerts(values []string) []*CryptoKeyCertificateChainsGooglePartitionCerts {
	var tValues []*CryptoKeyCertificateChainsGooglePartitionCerts
	for _, v := range values {
		tValues = append(tValues, &CryptoKeyCertificateChainsGooglePartitionCerts{
			Value: v,
		})
	}
	return tValues
}

var CryptoKeyTables = []interface{}{
	&CryptoKey{},
	&CryptoKeyCertificateChainsCaviumCerts{},
	&CryptoKeyCertificateChainsGoogleCardCerts{},
	&CryptoKeyCertificateChainsGooglePartitionCerts{},
}

func (c *Client) CryptoKeys(_ interface{}) error {
	c.log.Info("fetching crypto keys", zap.String("project", c.projectID))
	c.db.Where("project_id", c.projectID).Delete(CryptoKeyTables...)
	locations, err := c.getAllKmsLocations()
	if err != nil {
		return err
	}
	for _, l := range locations {
		c.log.Debug("fetching key rings for location", zap.String("location", l.Name))
		keys, err := c.getLocationKeyRings(l)
		if err != nil {
			return err
		}
		for _, k := range keys {
			c.log.Debug("fetching crypto keys of key ring", zap.String("key", k.Name))
			if err := c.getCryptoKeys(k); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) getCryptoKeys(key *cloudkms.KeyRing) error {
	nextPageToken := ""
	for {
		call := c.svc.Projects.Locations.KeyRings.CryptoKeys.List(key.Name)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(c.transformCryptoKeys(output.CryptoKeys))
		c.log.Info("populating CryptoKeys", zap.Int("count", len(output.CryptoKeys)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}


func (c *Client) getLocationKeyRings(location *cloudkms.Location) ([]*cloudkms.KeyRing, error) {
	var keys []*cloudkms.KeyRing
	call := c.svc.Projects.Locations.KeyRings.List(location.Name)
	nextPageToken := ""
	for {
		call.PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return nil, err
		}
		keys = append(keys, resp.KeyRings...)

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return keys, nil
}


func (c *Client) getAllKmsLocations() ([]*cloudkms.Location, error) {

	var locations []*cloudkms.Location
	call := c.svc.Projects.Locations.List("projects/" + c.projectID)
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
