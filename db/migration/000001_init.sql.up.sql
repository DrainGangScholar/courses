CREATE TABLE "school" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "type" varchar NOT NULL,
  "created_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "student" (
  "id" bigserial PRIMARY KEY,
  "school_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "enrollment" (
  "id" bigserial PRIMARY KEY,
  "student_id" bigint NOT NULL,
  "course_id" bigint NOT NULL,
  "created_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "course" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "attendants" bigint NOT NULL DEFAULT 0,
  "start_date" timestamptz NOT NULL DEFAULT (now()),
  "end_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "school" ("id");

CREATE INDEX ON "student" ("id");

CREATE INDEX ON "student" ("school_id");

CREATE INDEX ON "enrollment" ("id");

CREATE INDEX ON "enrollment" ("student_id");

CREATE INDEX ON "enrollment" ("course_id");

CREATE INDEX ON "enrollment" ("created_date");

CREATE INDEX ON "course" ("id");

CREATE INDEX ON "course" ("name");

CREATE INDEX ON "course" ("start_date");

CREATE INDEX ON "course" ("end_date");

ALTER TABLE "student" ADD FOREIGN KEY ("school_id") REFERENCES "school" ("id");

ALTER TABLE "enrollment" ADD FOREIGN KEY ("student_id") REFERENCES "student" ("id");

ALTER TABLE "enrollment" ADD FOREIGN KEY ("course_id") REFERENCES "course" ("id");
