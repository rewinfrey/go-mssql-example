-- CREATE DATABASE example

-- USE example
DROP TABLE IF EXISTS users

CREATE TABLE
  users (
    id BIGINT IDENTITY NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    age INT NOT NULL )