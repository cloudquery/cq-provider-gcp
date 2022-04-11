
# Table: gcp_compute_instance_group_instances

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_group_cq_id|uuid|Unique CloudQuery ID of gcp_compute_instance_groups table (FK)|
|instance|text|The URL of the instance|
|named_ports|jsonb|The named ports that belong to this instance group|
|status|text|"DEPROVISIONING" - The Nanny is halted and we are performing tear down tasks like network deprogramming, releasing quota, IP, tearing down disks etc   "PROVISIONING" - Resources are being allocated for the instance   "REPAIRING" - The instance is in repair   "RUNNING" - The instance is running   "STAGING" - All required resources have been allocated and the instance is being started   "STOPPED" - The instance has stopped successfully   "STOPPING" - The instance is currently stopping (either being deleted or killed)   "SUSPENDED" - The instance has suspended   "SUSPENDING" - The instance is suspending   "TERMINATED" - The instance has stopped (either by explicit action or underlying failure)|
