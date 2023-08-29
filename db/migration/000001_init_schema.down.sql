-- Drop foreign key constraints from cartDetails
ALTER TABLE IF EXISTS "cartDetails" DROP CONSTRAINT IF EXISTS "cartDetails_cart_id_fkey";
ALTER TABLE IF EXISTS "cartDetails" DROP CONSTRAINT IF EXISTS "cartDetails_product_id_fkey";

-- Drop foreign key constraints from carts
ALTER TABLE IF EXISTS "carts" DROP CONSTRAINT IF EXISTS "carts_username_fkey";

-- Drop foreign key constraints from orderDetails
ALTER TABLE IF EXISTS "orderDetails" DROP CONSTRAINT IF EXISTS "orderDetails_username_fkey";
ALTER TABLE IF EXISTS "orderDetails" DROP CONSTRAINT IF EXISTS "orderDetails_product_id_fkey";

-- Drop foreign key constraints from orders
ALTER TABLE IF EXISTS "orders" DROP CONSTRAINT IF EXISTS "orders_username_fkey";

-- Drop tables in reverse order of creation, using CASCADE to handle any missed dependencies
DROP TABLE IF EXISTS "cartDetails" CASCADE;
DROP TABLE IF EXISTS "carts" CASCADE;
DROP TABLE IF EXISTS "orderDetails" CASCADE;
DROP TABLE IF EXISTS "orders" CASCADE;
DROP TABLE IF EXISTS "products" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;
