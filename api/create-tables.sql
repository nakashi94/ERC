-- uuidを生成する拡張機能を導入
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists users(
  id uuid default uuid_generate_v4() not null,
  name varchar(32) not null,
  email varchar(255) not null,
  password varchar(255) not null,
  remenber_token varchar(255),
  refresh_token varchar(255),
  created_at timestamp default current_timestamp not null,
  updated_at timestamp,
  primary key (id)
);

create table if not exists recipes (
  id uuid default uuid_generate_v4() not null,
  user_id uuid not null,
  dish_name varchar(32) not null,
  url varchar(255) not null,
  category varchar(32) not null,
  media varchar(32) not null,
  repeat char(1) not null,
  cooking_time integer not null,
  created_at timestamp default current_timestamp not null,
  updated_at timestamp,
  primary key (id),
  foreign key (user_id) references users(id)
)
