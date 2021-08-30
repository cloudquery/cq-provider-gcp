
# Table: gcp_compute_url_map_header_action_request_headers_to_adds
Specification determining how headers are added to requests or responses
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|url_map_cq_id|uuid|Unique CloudQuery ID of gcp_compute_url_maps table (FK)|
|header_name|text|The name of the header|
|header_value|text|The value of the header to add|
|replace|boolean|If false, headerValue is appended to any values that already exist for the header If true, headerValue is set for the header, discarding any values that were set for that header The default value is false|
