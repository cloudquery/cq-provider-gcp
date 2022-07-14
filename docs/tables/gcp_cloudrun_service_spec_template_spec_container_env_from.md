
# Table: gcp_cloudrun_service_spec_template_spec_container_env_from
Not supported by Cloud Run EnvFromSource represents the source of a set of ConfigMaps
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_spec_template_spec_container_cq_id|uuid|Unique CloudQuery ID of gcp_cloudrun_service_spec_template_spec_containers table (FK)|
|config_map_ref_local_object_reference_name|text|Name of the referent|
|config_map_ref_name|text|The ConfigMap to select from|
|config_map_ref_optional|boolean|Specify whether the ConfigMap must be defined|
|prefix|text|An optional identifier to prepend to each key in the ConfigMap|
|secret_ref_local_object_reference_name|text|Name of the referent|
|secret_ref_name|text|The Secret to select from|
|secret_ref_optional|boolean|Specify whether the Secret must be defined|
