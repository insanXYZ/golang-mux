create table contacts
(
    id int primary key auto_increment,
    name varchar(25) not null ,
    email varchar(25) not null ,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp
) engine = innodb;