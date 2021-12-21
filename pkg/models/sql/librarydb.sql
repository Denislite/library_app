CREATE DATABASE librarydb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE librarydb;

CREATE TABLE authors (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    surname VARCHAR(30) NOT NULL,
    name VARCHAR(30) NOT NULL,
    middle_name VARCHAR(30) NOT NULL,
    image_path VARCHAR(100)
);

CREATE TABLE books (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL,
    alt_name VARCHAR(30),
    genre INTEGER NOT NULL,
    price INTEGER NOT NULL,
    count INTEGER NOT NULL,
    image_path VARCHAR(100),
    price_per_day INTEGER NOT NULL,
    year SMALLINT NOT NULL,
    reg_date DATE NOT NULL,
    rating SMALLINT
);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    surname VARCHAR(30) NOT NULL,
    name VARCHAR(30) NOT NULL,
    middle_name VARCHAR(30) NOT NULL,
    passport_data VARCHAR(30),
    birthday_date DATE NOT NULL,
    email VARCHAR(50) NOT NULL,
    address VARCHAR(255) NOT NULL,
    duty_count INTEGER NOT NULL DEFAULT 0,
    book_count INTEGER NOT NULL DEFAULT 0
);

CREATE Table genres (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL
);