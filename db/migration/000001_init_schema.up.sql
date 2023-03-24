CREATE TABLE "customers" (
                             "id" uuid PRIMARY KEY,
                             "username" varchar NOT NULL,
                             "hashed_password" varchar NOT NULL,
                             "full_name" varchar NOT NULL,
                             "email" varchar NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "drivers" (
                           "id" uuid PRIMARY KEY,
                           "username" varchar NOT NULL,
                           "full_name" varchar NOT NULL,
                           "cab_id" uuid NOT NULL,
                           "email" varchar NOT NULL,
                           "dob" date NOT NULL,
                           "location" point NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ratings" (
                           "id" uuid PRIMARY KEY,
                           "customer_id" uuid NOT NULL,
                           "driver_id" uuid NOT NULL,
                           "trip_id" uuid NOT NULL,
                           "rating" bigint NOT NULL,
                           "feedback" text
);

CREATE TABLE "trips" (
                         "id" uuid PRIMARY KEY,
                         "customer_id" uuid NOT NULL,
                         "driver_id" uuid,
                         "status" varchar NOT NULL,
                         "source" point NOT NULL,
                         "destination" point NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "customers" ("username");

CREATE UNIQUE INDEX ON "customers" ("username", "email");

CREATE INDEX ON "drivers" ("username");

CREATE UNIQUE INDEX ON "drivers" ("username", "email");

COMMENT ON COLUMN "ratings"."rating" IS 'must be positive';

ALTER TABLE "ratings" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id");
