CREATE TABLE IF NOT EXISTS products (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  price FLOAT NOT NULL,
  description VARCHAR(500) NOT NULL
);