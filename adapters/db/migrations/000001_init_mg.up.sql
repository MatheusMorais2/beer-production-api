CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "brewery" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar,
  "cnpj" varchar
);

CREATE TABLE "users" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar,
  "email" varchar unique,
  "password" varchar
);

CREATE TABLE "brewery_user" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "user_id" uuid,
  "brewery_id" uuid,
  "role" varchar
);

CREATE TABLE "recipe" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar,
  "brewery_id" uuid
);

CREATE TABLE "batch" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar,
  "recipe_id" uuid,
  "start_date" date,
  "finish_date" date
);

CREATE TABLE "recipe_step" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar,
  "recipe_id" uuid,
  "instruction" TEXT
  --"equipament_id" uuid
);

CREATE TABLE "batch_recipe_step" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "user_id" uuid,
  "recipe_step_id" uuid,
  "batch_id" uuid,
  "started_at" timestamp,
  "finished_at" timestamp
);

--CREATE TABLE "equipament" (
--  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
--  "name" varchar
--);

--CREATE TABLE "recipe_equipament" (
--  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
 -- "recipe_id" integer,
 -- "equipament_id" integer
--);

--ALTER TABLE "recipe_equipament" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("id");

--ALTER TABLE "recipe_equipament" ADD FOREIGN KEY ("equipament_id") REFERENCES "equipament" ("id");

--ALTER TABLE "recipe_step" ADD FOREIGN KEY ("equipament_id") REFERENCES "equipament" ("id");

ALTER TABLE "batch" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("id");

ALTER TABLE "recipe_step" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("id");

ALTER TABLE "batch_recipe_step" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "batch_recipe_step" ADD FOREIGN KEY ("recipe_step_id") REFERENCES "recipe_step" ("id");

ALTER TABLE "batch_recipe_step" ADD FOREIGN KEY ("batch_id") REFERENCES "batch" ("id");

ALTER TABLE "recipe" ADD FOREIGN KEY ("brewery_id") REFERENCES "brewery" ("id");

ALTER TABLE "brewery_user" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "brewery_user" ADD FOREIGN KEY ("brewery_id") REFERENCES "brewery" ("id");