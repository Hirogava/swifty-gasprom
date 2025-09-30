CREATE TYPE IF NOT EXISTS "statuses" AS ENUM (
  'expired',
  'pending',
  'claimed'
);

CREATE TYPE IF NOT EXISTS "bank_product_type" AS ENUM (
  'credit_card',
  'debit_card'
);

CREATE TYPE IF NOT EXISTS "bonus_type" AS ENUM (
  'cashback',
  'discount'
);

CREATE TABLE "bank_products" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "url" text NOT NULL,
  "type" bank_product_type NOT NULL
);

CREATE TABLE "user_used_bank_products" (
  "started_at" timestamp NOT NULL DEFAULT (now()),
  "bank_product_id" integer NOT NULL,
  "user_id" UUID NOT NULL
);

CREATE TABLE "bank_bonuses" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "url" text NOT NULL,
  "type" bonus_type NOT NULL
);

CREATE TABLE "user_bonuses" (
  "claimed_at" timestamp NOT NULL DEFAULT (now()),
  "status" statuses NOT NULL,
  "user_id" UUID NOT NULL,
  "bonus_id" integer NOT NULL
);

ALTER TABLE "user_used_bank_products" ADD FOREIGN KEY ("bank_product_id") REFERENCES "bank_products" ("id");

ALTER TABLE "user_used_bank_products" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_bonuses" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_bonuses" ADD FOREIGN KEY ("bonus_id") REFERENCES "bank_bonuses" ("id");
