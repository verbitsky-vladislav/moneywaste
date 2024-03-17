create table "User" (
    id serial primary key,
    nickname varchar(255) unique,
    password varchar(255)
)

