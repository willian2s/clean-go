CREATE TABLE product (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  price FLOAT NOT NULL,
  description VARCHAR(500) NOTNULL
);