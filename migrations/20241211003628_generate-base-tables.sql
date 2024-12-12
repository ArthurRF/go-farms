-- Create "farms" table
CREATE TABLE "public"."farms" (
	"id" bigserial NOT NULL,
	"farm_name" text NOT NULL,
	"land_area" bigint NOT NULL,
	"unit_of_measure" text NOT NULL,
	"address" text NOT NULL,
	"created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id")
);

-- Create "crops" table
CREATE TABLE "public"."crops" (
	"id" bigserial NOT NULL,
	"farm_id" bigint NOT NULL,
	"crop_type" text NOT NULL,
	"is_irrigated" boolean NOT NULL,
	"is_insured" boolean NOT NULL,
	"created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	CONSTRAINT "fk_farms_crops" FOREIGN KEY ("farm_id") REFERENCES "public"."farms" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);