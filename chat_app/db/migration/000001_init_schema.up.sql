CREATE TYPE "conversation_member_status" AS ENUM (
  'active',
  'deactive'
);

CREATE TABLE "conversations" (
  "conversation_id" varchar PRIMARY KEY NOT NULL,
  "name" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "conversation_members" (
  "user_id" varchar NOT NULL,
  "conversation_id" varchar NOT NULL,
  "status" conversation_member_status DEFAULT ('active'),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "conversation_id")
);

CREATE TABLE "users" (
  "user_id" varchar PRIMARY KEY NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "conversation_members" ADD FOREIGN KEY ("conversation_id") REFERENCES "conversations" ("conversation_id");

ALTER TABLE "conversation_members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
