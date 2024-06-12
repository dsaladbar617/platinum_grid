CREATE TABLE "sheets" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "short_name" varchar NOT NULL,
    "templates" varchar ,
);

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "man_number" NUMBER,
    "picture" VARCHAR,
    "email" VARCHAR(255) NOT NULL,
);

CREATE TABLE "user_roles" (
    "id" bigserial PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "role_name" VARCHAR(255) NOT NULL,
    "sheet_id" INTEGER NOT NULL,
    -- FOREIGN KEY (user_id) REFERENCES "users" (id),
    -- FOREIGN KEY (sheet_id) REFERENCES "sheets" (id)
);

CREATE TABLE "fields" (
    "id" bigserial PRIMARY KEY,
    "type" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "archived" BOOLEAN NOT NULL,
    "favorite" BOOLEAN NOT NULL,
    "sheet_id" INTEGER NOT NULL,
    -- FOREIGN KEY (sheet_id) REFERENCES "sheets" (id)
);

CREATE TABLE "entries"(
  "id" bigserial PRIMARY KEY,
  "sheet_id" INTEGER NOT NULL,
  "archived" BOOLEAN NOT NULL,
  -- FOREIGN KEY (sheet_id) REFERENCES "sheets" (id)
)

CREATE TABLE "values"(
  "id" bigserial PRIMARY KEY,
  "value" TEXT NOT NULL,
  "entry_id" INTEGER NOT NULL,
  "field_id" INTEGER NOT NULL,
  "value" TEXT NOT NULL,
  "checked" BOOLEAN NOT NULL DEFAULT FALSE,
  -- FOREIGN KEY (entry_id) REFERENCES "entries" (id),
  -- FOREIGN KEY (field_id) REFERENCES "fields" (id)

);

ALTER TABLE "user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_roles" ADD FOREIGN KEY ("sheet_id") REFERENCES "sheets" ("id");
ALTER TABLE "fields" ADD FOREIGN KEY ("sheet_id") REFERENCES "sheets" ("id");
ALTER TABLE "entries" ADD FOREIGN KEY ("sheet_id") REFERENCES "sheets" ("id");
ALTER TABLE "values" ADD FOREIGN KEY ("entry_id") REFERENCES "entries" ("id");
ALTER TABLE "values" ADD FOREIGN KEY ("field_id") REFERENCES "fields" ("id");