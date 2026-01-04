create database if not exists victoria_property;

use victoria_property;

create table if not exists roles (
    id int auto_increment primary key,
    name varchar(255) not null
);

create table if not exists users (
    id int auto_increment primary key,
    username varchar(255) not null,
    email varchar(255) UNIQUE not null,
    password varchar(255) not null,
    name varchar(255) not null,
    role_id int not null,

    foreign key (role_id) references roles(id)
);
