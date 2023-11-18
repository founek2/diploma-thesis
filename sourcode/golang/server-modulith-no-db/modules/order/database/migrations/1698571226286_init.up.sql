CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--bun:split

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
  FOREIGN KEY (order_seq_id) REFERENCES "order"(id)
)
--bun:split
CREATE UNIQUE INDEX uidx_rel_item_order ON rel_item_order (item_seq_id, order_seq_id); 
