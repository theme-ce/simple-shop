ALTER TABLE "orderDetails" ADD COLUMN "order_id" bigint;
ALTER TABLE "orderDetails" ADD CONSTRAINT "fk_orderDetails_order_id" FOREIGN KEY ("order_id") REFERENCES "orders"("id");
UPDATE "orderDetails" SET "order_id" = 2;
ALTER TABLE "orderDetails" ALTER COLUMN "order_id" SET NOT NULL;
