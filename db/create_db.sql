-- CREATE DATABASE IF NOT EXISTS server_db;
-- USE server_db;

CREATE TABLE IF NOT EXISTS 'users' (
  'uuid' blob PRIMARY KEY,
  'login' varchar(50) NOT NULL,
  'registration_date' date NOT NULL
);
