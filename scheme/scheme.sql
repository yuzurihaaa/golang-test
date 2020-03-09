drop table if exists users;

create table users (
    id serial not null primary key,
    name text not null
);

insert into users (name) values ('Yusuf');
insert into users (name) values ('Shira');
