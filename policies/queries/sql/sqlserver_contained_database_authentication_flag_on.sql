-- SELECT gsi.project_id, gsi.name, gsi.self_link AS link
-- FROM gcp_sql_instances gsi
-- WHERE database_version LIKE 'SQLSERVER%'
--     AND (settings_database_flags IS NULL
--         OR settings_database_flags ->> 'contained database authentication' != 'off'
--         OR settings_database_flags ->> 'contained database authentication' IS NULL);


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT gsi.name                                                                                                                                       AS resource_id,
       :'execution_time'::timestamp                                                                                                                   AS execution_time,
       :'framework'                                                                                                                                   AS framework,
       :'check_id'                                                                                                                                    AS check_id,
       'Ensure that the "contained database authentication" database flag for Cloud SQL on the SQL Server instance is set to "off" (Automated)' AS title,
       gsi.project_id                                                                                                                                 AS project_id,
       CASE
           WHEN
                       gsi.database_version LIKE 'SQLSERVER%'
                   AND (gsi.settings_database_flags IS NULL
                   OR gsi.settings_database_flags ->> 'contained database authentication' != 'off'
                   OR gsi.settings_database_flags ->> 'contained database authentication' IS NULL)
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                                                        AS status
FROM gcp_sql_instances gsi;