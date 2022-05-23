
# Table: gcp_container_cluster_autoscaling_resource_limits
Contains information about amount of some resource in the cluster
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of gcp_container_clusters table (FK)|
|maximum|bigint|Maximum amount of the resource in the cluster|
|minimum|bigint|Minimum amount of the resource in the cluster|
|resource_type|text|Resource name "cpu", "memory" or gpu-specific string|
