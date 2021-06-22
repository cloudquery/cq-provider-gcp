
# Table: gcp_kms_keyring_crypto_keys
CryptoKey: A CryptoKey represents a logical key that can be used for cryptographic operations
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|keyring_id|uuid|Unique ID of gcp_kms_keyrings table (FK)|
|project_id|text|GCP Project Id of the resource|
|location|text||
|policy|jsonb||
|create_time|timestamp without time zone|CreateTime: Output only|
|labels|jsonb|Labels: Labels with user-defined metadata|
|name|text|Name: Output only|
|next_rotation_time|timestamp without time zone|NextRotationTime: At next_rotation_time, the Key Management Service will automatically: 1|
|primary_algorithm|text|Algorithm: Output only|
|primary_attestation_cert_chains_cavium_certs|text[]|CaviumCerts: Cavium certificate chain corresponding to the attestation.|
|primary_attestation_cert_chains_google_card_certs|text[]|GoogleCardCerts: Google card certificate chain corresponding to the attestation.|
|primary_attestation_cert_chains_google_partition_certs|text[]|GooglePartitionCerts: Google partition certificate chain corresponding to the attestation.|
|primary_attestation_content|text|Content: Output only|
|primary_attestation_format|text|Format: Output only|
|primary_create_time|timestamp without time zone|CreateTime: Output only|
|primary_destroy_event_time|timestamp without time zone|DestroyEventTime: Output only|
|primary_destroy_time|timestamp without time zone|DestroyTime: Output only|
|primary_external_protection_level_options_external_key_uri|text|ExternalKeyUri: The URI for an external resource that this CryptoKeyVersion represents.|
|primary_generate_time|timestamp without time zone|GenerateTime: Output only|
|primary_import_failure_reason|text|ImportFailureReason: Output only|
|primary_import_job|text|ImportJob: Output only|
|primary_import_time|timestamp without time zone|ImportTime: Output only|
|primary_name|text|Name: Output only|
|primary_protection_level|text|ProtectionLevel: Output only|
|primary_state|text|State: The current state of the CryptoKeyVersion.  Possible values:   "CRYPTO_KEY_VERSION_STATE_UNSPECIFIED" - Not specified.   "PENDING_GENERATION" - This version is still being generated|
|purpose|text|Purpose: Immutable|
|rotation_period|text|RotationPeriod: next_rotation_time will be advanced by this period when the service automatically rotates a key|
|version_template_algorithm|text|Algorithm: Required|
|version_template_protection_level|text|ProtectionLevel: ProtectionLevel to use when creating a CryptoKeyVersion based on this template|
