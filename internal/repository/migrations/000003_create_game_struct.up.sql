CREATE TYPE risk_type AS ENUM (
  'crypto',
  'stocks',
  'bets',
  'questionable_projects'
);

CREATE TYPE career_category AS ENUM (
  'it_and_technology',
  'engineering_and_manufacturing',
  'medicine_and_healthcare',
  'marketing_and_sales',
  'working_professions'
);

CREATE TYPE news_type AS ENUM (
  'economic',
  'political',
  'corporate',
  'useless'
);

CREATE TYPE user_statuses AS ENUM (
  'finished',
  'active',
  'burnout',
  'bankrupt'
);

CREATE TYPE life_market_category AS ENUM (
  'entertainment_and_recreation',
  'appliances_and_gadgets',
  'gifts_and_social_interaction',
  'education_and_self_development',
  'health_and_care',
  'everyday_joys'
);

CREATE TABLE configs (
  id serial PRIMARY KEY,
  month_inflation decimal(10,2),
  month_dividend decimal(10,2)
);

CREATE TABLE careers (
  id serial PRIMARY KEY,
  name varchar(100),
  career_type career_category,
  career_level integer,
  career_grade integer,
  min_salary decimal(10,2),
  max_salary decimal(10,2),
  happiness_penalty_factor integer,
  natural_expenses decimal(10,2),
  months_to_grade integer
);

CREATE TABLE user_career (
  taked_at timestamp NOT NULL DEFAULT now(),
  career_level integer,
  career_grade integer,
  salary decimal(10,2) NOT NULL,
  career_id integer NOT NULL,
  natural_expenses decimal(10,2),
  user_id UUID NOT NULL
);

CREATE TABLE scenarios (
  id serial PRIMARY KEY,
  version integer NOT NULL,
  title varchar(150),
  type risk_type,
  crypto_category integer,
  min_price decimal(10,2),
  max_price decimal(10,2),
  start_win_chance integer,
  start_lose_chance integer,
  max_win decimal(10,2),
  max_lose decimal(10,2),
  start_price decimal(10,2)
);

CREATE TABLE risk (
  id serial PRIMARY KEY,
  user_id UUID NOT NULL,
  scenario_id integer NOT NULL,
  name varchar(100),
  type risk_type NOT NULL,
  price decimal(10,2),
  win_chance integer,
  lose_chance integer
);

CREATE TABLE crypto (
  id serial PRIMARY KEY,
  name varchar(255)
);

CREATE TABLE user_crypto_scenarios (
  id serial PRIMARY KEY,
  name varchar(150),
  user_id UUID NOT NULL,
  scenario_id integer NOT NULL,
  crypto_category integer NOT NULL,
  type risk_type,
  price decimal(10,2),
  win_chance integer,
  lose_chance integer,
  month integer
);

CREATE TABLE news (
  id serial PRIMARY KEY,
  type news_type NOT NULL,
  crypto_category integer,
  news_title varchar(255) NOT NULL,
  news_text text NOT NULL,
  effect_on_market integer,
  effect bool,
  version int
);

CREATE TABLE current_progress_news (
  month integer DEFAULT 1,
  news_id integer NOT NULL,
  user_id UUID NOT NULL
);

CREATE TABLE user_progress (
  id serial PRIMARY KEY,
  user_id UUID NOT NULL,
  money decimal(10,2) DEFAULT 0,
  happiness integer DEFAULT 100,
  health integer DEFAULT 100,
  month integer DEFAULT 1,
  status user_statuses NOT NULL,
  natural_expenses integer,
  last_sync_at timestamp
);

CREATE TABLE user_assets (
  user_id UUID NOT NULL,
  market_id integer NOT NULL
);

CREATE TABLE life_market (
  id serial PRIMARY KEY,
  name varchar(100),
  cost decimal(10,2),
  category life_market_category,
  effect_duration integer
);

CREATE TABLE events (
  id serial PRIMARY KEY,
  name varchar(255),
  event_text text,
  capital_percent integer,
  to_agree integer,
  refuse integer
);

CREATE TABLE user_buyed_risks (
  user_id UUID NOT NULL,
  risk_id integer,
  crypto_id integer
);

COMMENT ON COLUMN user_career.career_level IS 'Must be 1 to 5';

COMMENT ON COLUMN user_progress.happiness IS 'Must be 0 to 100';

COMMENT ON COLUMN user_progress.health IS 'Must be 0 to 100';

COMMENT ON COLUMN user_progress.month IS 'Must be 1 to 60';

COMMENT ON COLUMN current_progress_news.month IS 'Must be 1 to 60';

ALTER TABLE user_career ADD FOREIGN KEY (career_id) REFERENCES careers (id);

ALTER TABLE user_career ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE scenarios ADD FOREIGN KEY (crypto_category) REFERENCES crypto (id);

ALTER TABLE risk ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE risk ADD FOREIGN KEY (scenario_id) REFERENCES scenarios (id);

ALTER TABLE news ADD FOREIGN KEY (crypto_category) REFERENCES crypto (id);

ALTER TABLE user_crypto_scenarios ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_crypto_scenarios ADD FOREIGN KEY (scenario_id) REFERENCES scenarios (id);

ALTER TABLE user_crypto_scenarios ADD FOREIGN KEY (crypto_category) REFERENCES crypto (id);

ALTER TABLE current_progress_news ADD FOREIGN KEY (news_id) REFERENCES news (id);

ALTER TABLE current_progress_news ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_progress ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_assets ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_assets ADD FOREIGN KEY (market_id) REFERENCES life_market (id);

ALTER TABLE user_buyed_risks ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_buyed_risks ADD FOREIGN KEY (risk_id) REFERENCES risk (id);

ALTER TABLE user_buyed_risks ADD FOREIGN KEY (crypto_id) REFERENCES user_crypto_scenarios (id);

CREATE INDEX ON careers (name);