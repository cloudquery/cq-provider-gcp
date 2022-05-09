-- Resource: compute.instance_groups
CREATE TABLE IF NOT EXISTS "gcp_compute_instance_groups"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "project_id"         text,
    "creation_timestamp" timestamp,
    "description"        text,
    "fingerprint"        text,
    "id"                 bigint,
    "kind"               text,
    "name"               text,
    "named_ports"        jsonb,
    "network"            text,
    "region"             text,
    "self_link"          text,
    "size"               bigint,
    "subnetwork"         text,
    "zone"               text,
    CONSTRAINT gcp_compute_instance_groups_pk PRIMARY KEY (project_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "gcp_compute_instance_group_instances"
(
    "cq_id"                uuid NOT NULL,
    "cq_meta"              jsonb,
    "instance_group_cq_id" uuid,
    "instance"             text,
    "named_ports"          jsonb,
    "status"               text,
    CONSTRAINT gcp_compute_instance_group_instances_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (instance_group_cq_id) REFERENCES gcp_compute_instance_groups (cq_id) ON DELETE CASCADE
);
