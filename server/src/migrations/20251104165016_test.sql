-- Create "messages" table
CREATE TABLE "messages" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "message_text" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_messages_deleted_at" to table: "messages"
CREATE INDEX "idx_messages_deleted_at" ON "messages" ("deleted_at");
