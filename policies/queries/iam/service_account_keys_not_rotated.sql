-- SELECT
--     project_id,
--     gisa.id AS "account_id",
--     gisak.name AS "key_name",
--     gisak.valid_after_time
-- FROM gcp_iam_service_accounts gisa
--     JOIN gcp_iam_service_account_keys gisak ON
--             gisa.cq_id = gisak.service_account_cq_id
-- WHERE gisa.email LIKE '%iam.gserviceaccount.com'
--     AND gisak.valid_after_time <=
--         (now() - INTERVAL '90' day) -- noqa


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gisa.id                                                                                                AS resource_id,
                :'execution_time'::timestamp                                                                           AS execution_time,
                :'framework'                                                                                           AS framework,
                :'check_id'                                                                                            AS check_id,
                'Ensure user-managed/external keys for service accounts are rotated every 90 days or less (Automated)' AS title,
                gisa.project_id                                                                                        AS project_id,
                CASE
                    WHEN
                                gisa.email LIKE '%iam.gserviceaccount.com'
                            AND gisak.valid_after_time <=
                                (now() - INTERVAL '90' day)
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                                AS status
FROM gcp_iam_service_accounts gisa
         JOIN gcp_iam_service_account_keys gisak ON
    gisa.cq_id = gisak.service_account_cq_id;