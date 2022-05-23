
# Table: gcp_container_cluster_node_pool_config_taints
key, value, and effect
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_node_pool_cq_id|uuid|Unique CloudQuery ID of gcp_container_cluster_node_pools table (FK)|
|effect|text|"EFFECT_UNSPECIFIED" - Not set   "NO_SCHEDULE" - NoSchedule   "PREFER_NO_SCHEDULE" - PreferNoSchedule   "NO_EXECUTE" - NoExecute|
|key|text|Key for taint|
|value|text|Value for taint|
