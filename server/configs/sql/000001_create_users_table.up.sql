-- enable uuid extension
create extension if not exists "uuid-ossp";

-- create new users table
create table if not exists users(
  id uuid default uuid_generate_v4(),
  login varchar (32) not null unique,
  displayName varchar (32),
  password varchar (128) not null,
  email varchar (300) not null unique,

  createdAt timestamp default now(),
  updatedAt timestamp default now(),

  primary key (id)
);

create index users_login_idx on users(login);
