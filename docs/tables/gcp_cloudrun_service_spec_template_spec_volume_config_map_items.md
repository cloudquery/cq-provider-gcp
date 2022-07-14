
# Table: gcp_cloudrun_service_spec_template_spec_volume_config_map_items
Maps a string key to a path within a volume
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_spec_template_spec_volume_cq_id|uuid|Unique CloudQuery ID of gcp_cloudrun_service_spec_template_spec_volumes table (FK)|
|key|text|The Cloud Secret Manager secret version|
|mode|bigint|Mode bits to use on this file, must be a value between 0000 and 0777|
|path|text|The relative path of the file to map the key to|
