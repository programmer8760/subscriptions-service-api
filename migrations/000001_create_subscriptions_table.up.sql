CREATE TABLE IF NOT EXISTS subscriptions(
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  price integer NOT NULL,
  user_id uuid NOT NULL,
  start_date bigint NOT NULL,
  end_date bigint
);
