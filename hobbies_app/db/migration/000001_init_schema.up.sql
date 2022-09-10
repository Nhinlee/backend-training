CREATE TABLE "skills" (
  "skill_id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "user_name" varchar NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "hobbies" (
  "hobbie_id" bigserial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "continuos_day_max" integer NOT NULL,
  "skill_id" bigint,
  "user_id" bigint NOT NULL
);

CREATE TABLE "hobbie_logs" (
  "user_id" bigserial NOT NULL,
  "hobbie_id" bigint NOT NULL,
  "date_time" timestamptz DEFAULT (now()),
  PRIMARY KEY ("user_id", "hobbie_id")
);

ALTER TABLE "skills" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "hobbies" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("skill_id");

ALTER TABLE "hobbies" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "hobbie_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "hobbie_logs" ADD FOREIGN KEY ("hobbie_id") REFERENCES "hobbies" ("hobbie_id");
