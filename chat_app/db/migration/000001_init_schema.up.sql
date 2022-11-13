CREATE TYPE "conversation_user_status" AS ENUM (
  'active',
  'deactive'
);

CREATE TABLE "conversations" (
  "conversation_id" varchar PRIMARY KEY NOT NULL,
  "conversation_name" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "conversation_users" (
  "user_id" varchar NOT NULL,
  "conversation_id" varchar NOT NULL,
  "status" conversation_user_status DEFAULT ('active'),
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

ALTER TABLE "conversation_users" ADD FOREIGN KEY ("conversation_id") REFERENCES "conversations" ("conversation_id");

ALTER TABLE "conversation_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

-- DUMP Data for testing

-- Users
INSERT INTO public.users(
	user_id, first_name, last_name, email, hashed_password)
	VALUES ('u1', 'Nhin', 'Lee', 'nhinlechi@gmail.com', 'wefwefwefwefweffwef');
INSERT INTO public.users(
	user_id, first_name, last_name, email, hashed_password)
	VALUES ('u2', 'Eddie', 'Sin', 'eddie@gmail.com', 'wefwefweffwfwefwefwv');
INSERT INTO public.users(
	user_id, first_name, last_name, email, hashed_password)
	VALUES ('u3', 'Jacky', 'Jan', 'jacky@gmail.com', 'wefwvwwbwergwefwef');

-- Conversations
INSERT INTO public.conversations(
	conversation_id, conversation_name)
	VALUES ('c1', 'Test conversation');

-- Conversation Users
INSERT INTO public.conversation_users(
	user_id, conversation_id, status)
	VALUES ('u1', 'c1', 'active');
INSERT INTO public.conversation_users(
	user_id, conversation_id, status)
	VALUES ('u2', 'c1', 'active');
INSERT INTO public.conversation_users(
	user_id, conversation_id, status)
	VALUES ('u3', 'c1', 'active');