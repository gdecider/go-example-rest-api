create table users (
    id int auto_increment,
    email varchar(100) not null,
    encrypted_password varchar(255) not null,
    constraint users_pk primary key (id)
)