create table if not exists Users (
    user_id serial primary key,
    username varchar(100) not null,
    pc_name varchar(100) not null,
    user_group varchar(100),
    phone_number varchar(15),
    cabinet varchar(15),
    discription text,
    birthdate date,
    created_at TIMESTAMP not null,
    modified_at TIMESTAMP

);
