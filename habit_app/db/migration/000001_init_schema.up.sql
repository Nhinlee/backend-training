CREATE TABLE "skills" (
  "skill_id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "habits" (
  "habit_id" bigserial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "continuos_day_max" integer NOT NULL,
  "skill_id" bigint,
  "user_id" bigint NOT NULL
);

CREATE TABLE "habit_logs" (
  "user_id" bigserial NOT NULL,
  "habit_id" bigint NOT NULL,
  "date_time" timestamptz DEFAULT (now()),
  PRIMARY KEY ("user_id", "habit_id")
);

ALTER TABLE "skills" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "habits" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("skill_id");

ALTER TABLE "habits" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "habit_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "habit_logs" ADD FOREIGN KEY ("habit_id") REFERENCES "habits" ("habit_id");