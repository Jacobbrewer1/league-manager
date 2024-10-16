create table player
(
    id         int auto_increment
        primary key,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    email      varchar(255) not null,
    dob        date         not null,
    updated_at datetime     not null,
    constraint players_email_uindex
        unique (email)
)
    with system versioning;

