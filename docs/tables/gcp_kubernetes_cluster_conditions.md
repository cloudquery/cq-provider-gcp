
# Table: gcp_kubernetes_cluster_conditions
StatusCondition describes why a cluster or a node pool has a certain status (eg, ERROR or DEGRADED)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of gcp_kubernetes_clusters table (FK)|
|canonical_code|text|499 Client Closed Request   "UNKNOWN" - Unknown error|
|code|text|"UNKNOWN" - UNKNOWN indicates a generic condition   "GCE_STOCKOUT" - GCE_STOCKOUT indicates that Google Compute Engine resources are temporarily unavailable   "GKE_SERVICE_ACCOUNT_DELETED" - GKE_SERVICE_ACCOUNT_DELETED indicates that the user deleted their robot service account   "GCE_QUOTA_EXCEEDED" - Google Compute Engine quota was exceeded   "SET_BY_OPERATOR" - Cluster state was manually changed by an SRE due to a system logic error   "CLOUD_KMS_KEY_ERROR" - Unable to perform an encrypt operation against the CloudKMS key used for etcd level encryption   "CA_EXPIRING" - Cluster CA is expiring soon|
|message|text|Human-friendly representation of the condition|
