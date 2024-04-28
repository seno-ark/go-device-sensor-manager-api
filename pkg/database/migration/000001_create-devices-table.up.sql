CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "devices" (
  "id"          uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "name"        VARCHAR(100) NOT NULL,
  "description" TEXT NOT NULL,
  "status"      VARCHAR(20) NOT NULL,
  "created_at"  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);