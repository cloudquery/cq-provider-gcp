
# Table: gcp_compute_target_ssl_proxies
Represents a Target SSL Proxy resource  A target SSL proxy is a component of a SSL Proxy load balancer Global forwarding rules reference a target SSL proxy, and the target proxy then references an external backend service For more information, read Using Target Proxies (== resource_for {$api_version}
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|resource_id|text|Original Id of the resource|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource Provide this property when you create the resource|
|kind|text|Type of the resource Always compute#targetSslProxy for target SSL proxies|
|name|text|Name of the resource Provided by the client when the resource is created The name must be 1-63 characters long, and comply with RFC1035 Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash|
|proxy_header|text|Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1 The default is NONE|
|self_link|text|Server-defined URL for the resource|
|service|text|URL to the BackendService resource|
|ssl_certificates|text[]|URLs to SslCertificate resources that are used to authenticate connections to Backends At least one SSL certificate must be specified Currently, you may specify up to 15 SSL certificates|
|ssl_policy|text|URL of SslPolicy resource that will be associated with the TargetSslProxy resource If not set, the TargetSslProxy resource will not have any SSL policy configured|
