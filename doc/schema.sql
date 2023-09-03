-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-09-03T11:52:13.994Z

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "address" varchar,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "is_admin" boolean NOT NULL DEFAULT false
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "description" varchar NOT NULL,
  "price" double NOT NULL,
  "stock_quantity" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "total_price" double NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orderDetails" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "username" varchar NOT NULL,
  "quantity_ordered" bigint NOT NULL,
  "price_at_time_of_order" double NOT NULL
);

CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL
);

CREATE TABLE "cartDetails" (
  "id" bigserial PRIMARY KEY,
  "cart_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "quantity_added" bigint NOT NULL
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expired_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "orders" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "orderDetails" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "orderDetails" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orderDetails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "carts" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "cartDetails" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "cartDetails" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
