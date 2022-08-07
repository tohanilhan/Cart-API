SET TIMEZONE="Europe/Istanbul";
create database pf_db encoding 'UTF8';

\c pf_db

create schema if not exists pf_schema;
grant usage on schema pf_schema to postgres;

--access db
--revoke CONNECT ON DATABASE ctip_pg_auth_db FROM PUBLIC;
GRANT  CONNECT ON DATABASE pf_db  TO postgres;

--access schema
-- revoke ALL     ON SCHEMA ctip_pg_auth_schema FROM PUBLIC;


--all creation queries goes here: START

CREATE TABLE IF NOT EXISTS pf_schema.products
(
    product_id               uuid primary key,
    product_name             varchar(200),
    price                    float,
    vat                      varchar(4),
    quantity                   int2
);

CREATE INDEX index_product_name_from_products on pf_schema.products (product_name);


 CREATE TABLE IF NOT EXISTS pf_schema.orders(
    order_id               uuid,
	user_id                 uuid,
	cart                   json,
	discount               float,
	total_price_with_discount float,
    total_price_without_discount float,
    time bigint,
    timestamp text
);

CREATE INDEX index_user_id_from_orders on pf_schema.orders (user_id);


INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('81c815c6-d4cc-4a3f-9161-41540196d907','Iphone 13 Pro Max',999,18,30);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('b8e8cbb3-cb75-4a33-a80f-9e1e1033ec81','Xiaomi Redmi Note 11 Pro Plus 5G',729.99,18,69);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('87234a95-b9c5-47cd-8249-37ada8207a8c','Samsung S22 Ultra',899,18,45);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('19ad5d05-2abd-41fa-9f7b-ef697b84a209','Playstation 5',999,8,20);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('2eabefa9-a9ca-409d-849d-52b41ec391c9','ASUS ROG STRIX 17.6 inch 120 Hertz 2 TB HDD 512 GB SSD RTX 3070 I7 11600K Gamer Laptop',1799,8,250);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('4d7d0011-77f2-4dfa-80c3-bfd67d9a312b','17.6 inch Laptop HardCase',54.99,1,30);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('ff5d8f75-cb9b-4dce-817c-9b0a3f1715c8','MacBook Pro',1999,18,178);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('32746f0a-28ff-4db9-99ed-cdba7c57e512','Huawei Backpack',59.99,1,100);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('f10737f0-0fb9-4202-9cc2-f39f2adc96da','Xbox Series X',799.99,8,100);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('dc689ad5-0c2b-4150-9b11-5fe96db09422','Airpods Pro',289.99,8,100);
INSERT INTO pf_schema.products(product_id, product_name, price, vat, quantity) VALUES ('43565584-4b2b-482f-9fc6-b947cac437c8','Samsung Galaxy Buds Pro',159.99,1,100);
