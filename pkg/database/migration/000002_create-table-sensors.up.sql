CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "sensors" (
  "id"          uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "device_id"   uuid NOT NULL,
  "type"        VARCHAR(50) NOT NULL,
  "name"        VARCHAR(100) NOT NULL,
  "description" TEXT NOT NULL,
  "created_at"  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);