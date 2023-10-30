CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--bun:split

CREATE TABLE payment (
  id SERIAL PRIMARY KEY,
  payment_id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  invoice_seq_id bigint,
  credit_card_number bigint,
  amount double precision,
  created_at timestamp default now()
)