CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--bun:split

CREATE TABLE item (
  id SERIAL PRIMARY KEY,
  item_id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  title text,
  description text,
  quantity bigint,
  price double precision,
  created_at timestamp default now()
);

CREATE TABLE "order" (
  id SERIAL PRIMARY KEY,
  order_id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  title text,
  description text,
  quantity bigint,
  price double precision,
  ship_date timestamp,
  status text,
  complete boolean,
  created_at timestamp default now()
);

--bun:split

CREATE TABLE rel_item_order (
  id SERIAL PRIMARY KEY,
  item_seq_id bigint,
  order_seq_id bigint,
  FOREIGN KEY (item_seq_id) REFERENCES item(id),
  FOREIGN KEY (order_seq_id) REFERENCES "order"(id)
)
--bun:split
CREATE UNIQUE INDEX uidx_rel_item_order ON rel_item_order (item_seq_id, order_seq_id); 

--bun:split

CREATE TABLE invoice (
  id SERIAL PRIMARY KEY,
  invoice_id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  order_seq_id bigint,
  price double precision,
  status text,
  pdf_link text,
  created_at timestamp default now(),
  FOREIGN KEY (order_seq_id) REFERENCES "order"(id)
)

--bun:split

CREATE TABLE payment (
  id SERIAL PRIMARY KEY,
  payment_id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  invoice_seq_id bigint,
  credit_card_number bigint,
  amount double precision,
  created_at timestamp default now(),
  FOREIGN KEY (invoice_seq_id) REFERENCES "invoice"(id)
)

--bun:split

CREATE TABLE cart (
  id SERIAL PRIMARY KEY,
  cart_id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  user_id text,
  total_quantity bigint DEFAULT 0,
  created_at timestamp default now()
)
--bun:split
CREATE UNIQUE INDEX uidx_cart_user_id ON cart (user_id); 

--bun:split

CREATE TABLE rel_item_cart (
  id SERIAL PRIMARY KEY,
  item_seq_id bigint,
  cart_seq_id bigint,
  FOREIGN KEY (item_seq_id) REFERENCES item(id),
  FOREIGN KEY (cart_seq_id) REFERENCES cart(id)
)
--bun:split
CREATE UNIQUE INDEX uidx_rel_item_cart ON rel_item_cart (item_seq_id, cart_seq_id); 