drop database if exists victoria_property;
create database if not exists victoria_property;

use victoria_property;

create table if not exists roles (
    id int auto_increment primary key,
    name varchar(10) not null
);

create table if not exists users (
    id int auto_increment primary key,
    email varchar(255) UNIQUE not null,
    password varchar(255) not null,
    username varchar(255) not null,
    phone_number varchar(20) not null,
    role_id int not null,

    foreign key (role_id) references roles(id)
);

create table if not exists property_types (
    id int auto_increment primary key,
    name varchar(20) not null
);

create table if not exists properties (
    id int auto_increment primary key,
    title varchar(255) not null,
    description text,
    price bigint not null,
    status int not null,
    province varchar(50) not null,
    regency varchar(50) not null,
    district varchar(50) not null,
    address varchar(100),
    building_area int,
    land_area int,
    electricity int,
    water_source varchar(50),
    bedrooms int,
    bathrooms int,
    floors int,
    garage int,
    carport int,
    certificate varchar(20),
    year_constructed int,
    sale_type varchar(20) not null,
    created_at timestamp default current_timestamp,
    cover_image_url varchar (100), 
    property_type_id int not null,
    agent_id int not null,

    foreign key (property_type_id) references property_types(id),
    foreign key (agent_id) references users(id)
);

create table if not exists property_images (
    id int auto_increment primary key,
    url varchar(255) not null,
    property_id int not null,

    foreign key (property_id) references properties(id) on delete cascade
);

create table if not exists favorites (
    id int auto_increment primary key,
    user_id int not null,
    property_id int not null,

    foreign key (user_id) references users(id) on delete cascade,
    foreign key (property_id) references properties(id) on delete cascade
);
