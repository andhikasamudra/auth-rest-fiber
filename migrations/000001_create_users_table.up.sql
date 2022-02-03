create extension "uuid-ossp";

CREATE TABLE users (
  id serial PRIMARY KEY,
  guid uuid NOT NULL DEFAULT uuid_generate_v4(),
  email varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now()
);