CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--bun:split

CREATE TABLE cart (
  id SERIAL PRIMARY KEY,
  cart_id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  user_id text,
  total_quantity bigint DEFAULT 0,
  created_at timestamp default now()
)

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