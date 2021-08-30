
# Table: gcp_compute_url_map_weighted_backend_services
In contrast to a single BackendService in HttpRouteAction to which all matching traffic is directed to, WeightedBackendService allows traffic to be split across multiple BackendServices
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|url_map_cq_id|uuid|Unique CloudQuery ID of gcp_compute_url_maps table (FK)|
|backend_service|text|The full or partial URL to the default BackendService resource Before forwarding the request to backendService, the loadbalancer applies any relevant headerActions specified as part of this backendServiceWeight|
|header_action_request_headers_to_remove|text[]|A list of header names for headers that need to be removed from the request prior to forwarding the request to the backendService|
|header_action_response_headers_to_remove|text[]|A list of header names for headers that need to be removed from the response prior to sending the response back to the client|
|weight|bigint|Specifies the fraction of traffic sent to backendService, computed as weight / (sum of all weightedBackendService weights in routeAction)  The selection of a backend service is determined only for new traffic Once a user's request has been directed to a backendService, subsequent requests will be sent to the same backendService as determined by the BackendService's session affinity policy|
