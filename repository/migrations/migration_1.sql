create table "User" (
    id serial primary key,
    nickname char(255) unique,
    password char(32)
)

a