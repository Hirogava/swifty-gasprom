CREATE TABLE "game_logs" (
  "id" serial PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "event_id" integer,
  "action" jsonb
);

CREATE TABLE "analytics_snapshots" (
  "id" serial PRIMARY KEY,
  "metric" jsonb NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "game_logs" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
