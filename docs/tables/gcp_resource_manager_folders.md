
# Table: gcp_resource_manager_folders
Folder: A folder in an organization's resource hierarchy, used to organize that organization's resources. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|policy|jsonb||
|create_time|timestamp without time zone|CreateTime: Output only|
|delete_time|timestamp without time zone|DeleteTime: Output only|
|display_name|text|DisplayName: The folder's display name|
|etag|text|Etag: Output only|
|name|text|Name: Output only|
|parent|text|Parent: Required|
|state|text|State: Output only|
|update_time|timestamp without time zone|UpdateTime: Output only|
