CREATE TABLE "skills" (
  "skill_id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL
);

CREATE TABLE "habits" (
  "habit_id" bigserial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "target_consecutive_days" integer NOT NULL,
  "skill_id" bigint,
  "user_id" bigint NOT NULL
);

CREATE TABLE "habit_logs" (
  "user_id" bigint NOT NULL,
  "habit_id" bigint NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("user_id", "habit_id", "created_at")
);

ALTER TABLE "skills" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "habits" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("skill_id");

ALTER TABLE "habits" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "habit_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "habit_logs" ADD FOREIGN KEY ("habit_id") REFERENCES "habits" ("habit_id");

-- Just for testing
INSERT INTO public.users(
	first_name, last_name, email, password)
	VALUES ('Nhin', 'Le Chi', 'nhinlechi@gmail.com', '123456');

-- Just for testing
INSERT INTO public.skills(
	user_id, title)
	VALUES (1, 'algorithm & data structure');
INSERT INTO public.skills(
	user_id, title)
	VALUES (1, 'english');
INSERT INTO public.skills(
	user_id, title)
	VALUES (1, 'learning');