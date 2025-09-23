CREATE TYPE "event_type" AS ENUM (
  'crypto',
  'stocks',
  'bets',
  'questionable_projects'
);

CREATE TYPE "news_type" AS ENUM (
  'economic',
  'political',
  'corporate',
  'useless'
);

CREATE TYPE "user_statuses" AS ENUM (
  'finished',
  'active',
  'burnout',
  'bankrupt'
);

CREATE TABLE "configs" (
  "id" serial PRIMARY KEY,
  "global_settings" jsonb NOT NULL
);

CREATE TABLE "careers" (
  "id" serial PRIMARY KEY,
  "field_id" integer NOT NULL,
  "name" varchar(150) NOT NULL,
  "salary_growth_factor" decimal(10,2),
  "happiness_penalty_factor" decimal(10,2)
);

CREATE TABLE "career_fields" (
  "id" serial PRIMARY KEY,
  "name" varchar(100)
);

CREATE TABLE "user_career" (
  "taked_at" timestamp NOT NULL DEFAULT (now()),
  "career_level" integer,
  "career_id" integer NOT NULL,
  "user_id" UUID NOT NULL
);

CREATE TABLE "scenarios" (
  "id" serial PRIMARY KEY,
  "version" integer NOT NULL,
  "title" varchar(150),
  "description" text
);

CREATE TABLE "events" (
  "id" serial PRIMARY KEY,
  "scenario_id" integer NOT NULL,
  "name" varchar(100),
  "type" event_type NOT NULL,
  "effect_money" decimal(10,2),
  "effect_happiness" integer,
  "effect_health" integer,
  "probability" decimal(3,2)
);

CREATE TABLE "news" (
  "id" serial PRIMARY KEY,
  "type" news_type NOT NULL,
  "news_title" varchar(255) NOT NULL,
  "news_text" text NOT NULL,
  "effect_on_market" jsonb,
  "version" int
);

CREATE TABLE "current_progress_news" (
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "news_id" integer NOT NULL,
  "user_id" UUID NOT NULL
);

CREATE TABLE "user_progress" (
  "id" serial PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "money" decimal(10,2) DEFAULT 0,
  "happiness" integer DEFAULT 100,
  "health" integer DEFAULT 100,
  "month" integer DEFAULT 1,
  "status" user_statuses NOT NULL,
  "last_sync_at" timestamp
);

CREATE TABLE "user_assets" (
  "user_id" UUID NOT NULL,
  "market_id" integer NOT NULL
);

CREATE TABLE "life_market" (
  "id" serial PRIMARY KEY,
  "name" varchar(100),
  "effect_duration" integer
);

COMMENT ON COLUMN "user_career"."career_level" IS 'Must be 1 to 5';

COMMENT ON COLUMN "user_progress"."happiness" IS 'Must be 0 to 100';

COMMENT ON COLUMN "user_progress"."health" IS 'Must be 0 to 100';

COMMENT ON COLUMN "user_progress"."month" IS 'Must be 1 to 60';

ALTER TABLE "careers" ADD FOREIGN KEY ("field_id") REFERENCES "career_fields" ("id");

ALTER TABLE "user_career" ADD FOREIGN KEY ("career_id") REFERENCES "careers" ("id");

ALTER TABLE "user_career" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("scenario_id") REFERENCES "scenarios" ("id");

ALTER TABLE "current_progress_news" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id");

ALTER TABLE "current_progress_news" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_progress" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_assets" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_assets" ADD FOREIGN KEY ("market_id") REFERENCES "life_market" ("id");
