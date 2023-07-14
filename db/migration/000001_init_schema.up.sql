CREATE TABLE "province" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL UNIQUE,
  "code" varchar NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now())

);

INSERT INTO province (name, code) VALUES
	('Alberta', 'AB'),
	('British Columbia', 'BC'),
	('Manitoba', 'MB'),
	('New Brunswick', 'NB'),
	('Newfoundland and Labrador', 'NL'),
	('Northwest Territories', 'NT'),
	('Nova Scotia', 'NS'),
	('Nunavut', 'NU'),
	('Ontario', 'ON'),
	('Prince Edward Island', 'PE'),
	('Quebec', 'QC'),
	('Saskatchewan', 'SK'),
	('Yukon', 'YT');

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "prov_id" int,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


ALTER TABLE "users" ADD FOREIGN KEY ("prov_id") REFERENCES "province" ("id");
