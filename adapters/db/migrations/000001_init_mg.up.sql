CREATE TABLE "brewery" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "cnpj" varchar
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "role" varchar,
  "brewery_id" integer
);

CREATE TABLE "recipe" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "brewery_id" integer
);

CREATE TABLE "equipament" (
  "id" integer PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "recipe_equipament" (
  "id" integer PRIMARY KEY,
  "recipe_id" integer,
  "equipament_id" integer
);

CREATE TABLE "batch" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "recipe_id" integer,
  "start_date" date,
  "finish_date" date
);

CREATE TABLE "recipe_step" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "recipe_id" integer,
  "equipament_id" integer
);

CREATE TABLE "batch_recipe_step" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "recipe_step_id" integer,
  "batch_id" integer,
  "started_at" timestamp,
  "finished_at" timestamp
);

ALTER TABLE "recipe_equipament" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("id");

ALTER TABLE "recipe_equipament" ADD FOREIGN KEY ("equipament_id") REFERENCES "equipament" ("id");

ALTER TABLE "batch" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("id");

ALTER TABLE "recipe_step" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("id");

ALTER TABLE "recipe_step" ADD FOREIGN KEY ("equipament_id") REFERENCES "equipament" ("id");

ALTER TABLE "batch_recipe_step" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "batch_recipe_step" ADD FOREIGN KEY ("recipe_step_id") REFERENCES "recipe_step" ("id");

ALTER TABLE "batch_recipe_step" ADD FOREIGN KEY ("batch_id") REFERENCES "batch" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("brewery_id") REFERENCES "brewery" ("id");

ALTER TABLE "recipe" ADD FOREIGN KEY ("brewery_id") REFERENCES "brewery" ("id");