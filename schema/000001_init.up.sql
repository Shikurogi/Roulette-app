CREATE TABLE IF NOT EXISTS users(
    id          serial             not null unique,
    name        varchar(255)       not null unique,
    age         varchar(255)       not null,
    password    varchar(255)       not null,
    token       varchar(255)
);