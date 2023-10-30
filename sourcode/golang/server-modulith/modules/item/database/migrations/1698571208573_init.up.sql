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

