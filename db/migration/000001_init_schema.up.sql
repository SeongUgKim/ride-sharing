CREATE TABLE "customers" (
                             "id" uuid PRIMARY KEY,
                             "username" varchar NOT NULL UNIQUE,
                             "hashed_password" varchar NOT NULL,
                             "full_name" varchar NOT NULL,
                             "email" varchar NOT NULL,
                             "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "drivers" (
                           "id" uuid PRIMARY KEY,
                           "username" varchar NOT NULL UNIQUE,
                           "hashed_password" varchar NOT NULL,
                           "full_name" varchar NOT NULL,
                           "email" varchar NOT NULL,
                           "cab_id" uuid NOT NULL,
                           "dob" date NOT NULL,
                           "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
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

CREATE TABLE "cabs" (
    "id" uuid PRIMARY KEY,
    "cab_type" varchar NOT NULL,
    "reg_no" uuid NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);


COMMENT ON COLUMN "ratings"."rating" IS 'must be positive';

ALTER TABLE "drivers" ADD FOREIGN KEY ("cab_id") REFERENCES "cabs" ("id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id");
