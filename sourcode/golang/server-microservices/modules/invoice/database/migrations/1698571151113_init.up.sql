CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--bun:split

CREATE TABLE invoice (
  id SERIAL PRIMARY KEY,
  invoice_id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  order_seq_id bigint,
  price double precision,
  status text,
  pdf_link text,
  created_at timestamp default now()
)
