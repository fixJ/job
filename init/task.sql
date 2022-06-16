create table task (
    id int auto_increment not null,
    created_at int not null,
    updated_at int not null,
    name varchar(128) not null,
    target varchar(32) not null,
    command varchar(256) not null,
    status int not null,
    cron boolean not null,
    runtime int not null,
    primary key (id)
)engine=InnoDB default character set utf8mb4