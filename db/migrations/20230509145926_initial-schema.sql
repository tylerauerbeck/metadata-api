-- +goose Up
-- create "annotation_namespaces" table
CREATE TABLE "annotation_namespaces" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "tenant_id" character varying NOT NULL, "private" boolean NOT NULL DEFAULT false, PRIMARY KEY ("id"));
-- create index "annotationnamespace_created_at" to table: "annotation_namespaces"
CREATE INDEX "annotationnamespace_created_at" ON "annotation_namespaces" ("created_at");
-- create index "annotationnamespace_tenant_id" to table: "annotation_namespaces"
CREATE INDEX "annotationnamespace_tenant_id" ON "annotation_namespaces" ("tenant_id");
-- create index "annotationnamespace_tenant_id_name" to table: "annotation_namespaces"
CREATE UNIQUE INDEX "annotationnamespace_tenant_id_name" ON "annotation_namespaces" ("tenant_id", "name");
-- create index "annotationnamespace_updated_at" to table: "annotation_namespaces"
CREATE INDEX "annotationnamespace_updated_at" ON "annotation_namespaces" ("updated_at");
-- create "metadata" table
CREATE TABLE "metadata" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "node_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "metadata_created_at" to table: "metadata"
CREATE INDEX "metadata_created_at" ON "metadata" ("created_at");
-- create index "metadata_node_id_key" to table: "metadata"
CREATE UNIQUE INDEX "metadata_node_id_key" ON "metadata" ("node_id");
-- create index "metadata_updated_at" to table: "metadata"
CREATE INDEX "metadata_updated_at" ON "metadata" ("updated_at");
-- create "annotations" table
CREATE TABLE "annotations" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "json_data" jsonb NOT NULL, "annotation_namespace_id" character varying NOT NULL, "metadata_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "annotations_annotation_namespaces_namespace" FOREIGN KEY ("annotation_namespace_id") REFERENCES "annotation_namespaces" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "annotations_metadata_metadata" FOREIGN KEY ("metadata_id") REFERENCES "metadata" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "annotation_annotation_namespace_id_json_data" to table: "annotations"
CREATE INDEX "annotation_annotation_namespace_id_json_data" ON "annotations" USING gin ("annotation_namespace_id", "json_data");
-- create index "annotation_created_at" to table: "annotations"
CREATE INDEX "annotation_created_at" ON "annotations" ("created_at");
-- create index "annotation_metadata_id_annotation_namespace_id" to table: "annotations"
CREATE UNIQUE INDEX "annotation_metadata_id_annotation_namespace_id" ON "annotations" ("metadata_id", "annotation_namespace_id");
-- create index "annotation_updated_at" to table: "annotations"
CREATE INDEX "annotation_updated_at" ON "annotations" ("updated_at");
-- create "status_namespaces" table
CREATE TABLE "status_namespaces" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "resource_provider_id" character varying NOT NULL, "private" boolean NOT NULL DEFAULT false, PRIMARY KEY ("id"));
-- create index "statusnamespace_created_at" to table: "status_namespaces"
CREATE INDEX "statusnamespace_created_at" ON "status_namespaces" ("created_at");
-- create index "statusnamespace_resource_provider_id" to table: "status_namespaces"
CREATE INDEX "statusnamespace_resource_provider_id" ON "status_namespaces" ("resource_provider_id");
-- create index "statusnamespace_resource_provider_id_name" to table: "status_namespaces"
CREATE UNIQUE INDEX "statusnamespace_resource_provider_id_name" ON "status_namespaces" ("resource_provider_id", "name");
-- create index "statusnamespace_updated_at" to table: "status_namespaces"
CREATE INDEX "statusnamespace_updated_at" ON "status_namespaces" ("updated_at");
-- create "status" table
CREATE TABLE "status" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "source" character varying NOT NULL, "json_data" jsonb NOT NULL, "status_namespace_id" character varying NOT NULL, "metadata_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "status_metadata_metadata" FOREIGN KEY ("metadata_id") REFERENCES "metadata" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "status_status_namespaces_namespace" FOREIGN KEY ("status_namespace_id") REFERENCES "status_namespaces" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "status_created_at" to table: "status"
CREATE INDEX "status_created_at" ON "status" ("created_at");
-- create index "status_metadata_id_status_namespace_id" to table: "status"
CREATE INDEX "status_metadata_id_status_namespace_id" ON "status" ("metadata_id", "status_namespace_id");
-- create index "status_metadata_id_status_namespace_id_source" to table: "status"
CREATE UNIQUE INDEX "status_metadata_id_status_namespace_id_source" ON "status" ("metadata_id", "status_namespace_id", "source");
-- create index "status_status_namespace_id_json_data" to table: "status"
CREATE INDEX "status_status_namespace_id_json_data" ON "status" USING gin ("status_namespace_id", "json_data");
-- create index "status_updated_at" to table: "status"
CREATE INDEX "status_updated_at" ON "status" ("updated_at");

-- +goose Down
-- reverse: create index "status_updated_at" to table: "status"
DROP INDEX "status_updated_at";
-- reverse: create index "status_status_namespace_id_json_data" to table: "status"
DROP INDEX "status_status_namespace_id_json_data";
-- reverse: create index "status_metadata_id_status_namespace_id_source" to table: "status"
DROP INDEX "status_metadata_id_status_namespace_id_source";
-- reverse: create index "status_metadata_id_status_namespace_id" to table: "status"
DROP INDEX "status_metadata_id_status_namespace_id";
-- reverse: create index "status_created_at" to table: "status"
DROP INDEX "status_created_at";
-- reverse: create "status" table
DROP TABLE "status";
-- reverse: create index "statusnamespace_updated_at" to table: "status_namespaces"
DROP INDEX "statusnamespace_updated_at";
-- reverse: create index "statusnamespace_resource_provider_id_name" to table: "status_namespaces"
DROP INDEX "statusnamespace_resource_provider_id_name";
-- reverse: create index "statusnamespace_resource_provider_id" to table: "status_namespaces"
DROP INDEX "statusnamespace_resource_provider_id";
-- reverse: create index "statusnamespace_created_at" to table: "status_namespaces"
DROP INDEX "statusnamespace_created_at";
-- reverse: create "status_namespaces" table
DROP TABLE "status_namespaces";
-- reverse: create index "annotation_updated_at" to table: "annotations"
DROP INDEX "annotation_updated_at";
-- reverse: create index "annotation_metadata_id_annotation_namespace_id" to table: "annotations"
DROP INDEX "annotation_metadata_id_annotation_namespace_id";
-- reverse: create index "annotation_created_at" to table: "annotations"
DROP INDEX "annotation_created_at";
-- reverse: create index "annotation_annotation_namespace_id_json_data" to table: "annotations"
DROP INDEX "annotation_annotation_namespace_id_json_data";
-- reverse: create "annotations" table
DROP TABLE "annotations";
-- reverse: create index "metadata_updated_at" to table: "metadata"
DROP INDEX "metadata_updated_at";
-- reverse: create index "metadata_node_id_key" to table: "metadata"
DROP INDEX "metadata_node_id_key";
-- reverse: create index "metadata_created_at" to table: "metadata"
DROP INDEX "metadata_created_at";
-- reverse: create "metadata" table
DROP TABLE "metadata";
-- reverse: create index "annotationnamespace_updated_at" to table: "annotation_namespaces"
DROP INDEX "annotationnamespace_updated_at";
-- reverse: create index "annotationnamespace_tenant_id_name" to table: "annotation_namespaces"
DROP INDEX "annotationnamespace_tenant_id_name";
-- reverse: create index "annotationnamespace_tenant_id" to table: "annotation_namespaces"
DROP INDEX "annotationnamespace_tenant_id";
-- reverse: create index "annotationnamespace_created_at" to table: "annotation_namespaces"
DROP INDEX "annotationnamespace_created_at";
-- reverse: create "annotation_namespaces" table
DROP TABLE "annotation_namespaces";
