CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "owner" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" VARCHAR NOT NULL,
    "token" VARCHAR NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "template" (
    "id" SERIAL NOT NULL,
    "owner_id" uuid NOT NULL,
    "name" VARCHAR NOT NULL,
    "html_body" TEXT NOT NULL DEFAULT '',
    "text_body" TEXT NOT NULL DEFAULT '',
    PRIMARY KEY ("id"),
    FOREIGN KEY ("owner_id") REFERENCES "owner" ("id")
);

CREATE TABLE "mailer" (
    "id" SERIAL NOT NULL,
    "owner_id" uuid NOT NULL,
    "name" VARCHAR NOT NULL,
    "token" VARCHAR,
    "host" VARCHAR,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("owner_id") REFERENCES "owner" ("id")
);

CREATE TABLE "message" (
    "id" SERIAL NOT NULL,
    "owner_id" uuid NOT NULL,
    "name" VARCHAR NOT NULL,
    "mail_from" VARCHAR NOT NULL,
    "sender_name" VARCHAR NOT NULL,
    "subject" VARCHAR NOT NULL DEFAULT '',
    "template_id" integer NOT NULL,
    "default_mailer_id" integer NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("owner_id") REFERENCES "owner" ("id"),
    FOREIGN KEY ("template_id") REFERENCES "template" ("id"),
    FOREIGN KEY ("default_mailer_id") REFERENCES "mailer" ("id")
);

CREATE INDEX ON "template" ("owner_id");
CREATE INDEX ON "mailer" ("owner_id");
CREATE INDEX ON "message" ("owner_id");
CREATE INDEX ON "message" ("template_id");