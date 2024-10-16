create table players
(
    id         int          not null auto_increment,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    email      varchar(255) not null,
    dob        date         not null,
    updated_at datetime     not null,
    primary key (id),
    constraint players_email_uindex
        unique (email)
);
