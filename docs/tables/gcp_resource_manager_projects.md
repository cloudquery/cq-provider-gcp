
# Table: gcp_resource_manager_projects
A project is a high-level Google Cloud entity It is a container for ACLs, APIs, App Engine Apps, VMs, and other Google Cloud Platform resources
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|policy|jsonb||
|create_time|timestamp without time zone|Creation time|
|delete_time|timestamp without time zone|The time at which this resource was requested for deletion|
|display_name|text|A user-assigned display name of the project When present it must be between 4 to 30 characters Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point|
|etag|text|A checksum computed by the server based on the current value of the Project resource This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding|
|labels|jsonb|The labels associated with this project Label keys must be between 1 and 63 characters long and must conform to the following regular expression: \a-z\ (\[-a-z0-9\]*\[a-z0-9\])? Label values must be between 0 and 63 characters long and must conform to the regular expression (\a-z\ (\[-a-z0-9\]*\[a-z0-9\])?)? No more than 256 labels can be associated with a given resource Clients should store labels in a representation such as JSON that does not depend on specific characters being disallowed|
|name|text|The unique resource name of the project It is an int64 generated number prefixed by "projects/"|
|parent|text|A reference to a parent Resource eg, `organizations/123` or `folders/876`|
|project_id|text|GCP Project Id of the resource|
|state|text|The project lifecycle state  Possible values:   "STATE_UNSPECIFIED" - Unspecified state This is only used/useful for distinguishing unset values   "ACTIVE" - The normal and active state   "DELETE_REQUESTED" - The project has been marked for deletion by the user (by invoking DeleteProject) or by the system (Google Cloud Platform) This can generally be reversed by invoking UndeleteProject|
|update_time|timestamp without time zone|The most recent time this resource was modified|
