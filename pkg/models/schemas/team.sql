create table teams
(
    id             int          not null auto_increment,
    name           varchar(25)  not null,
    contact_email  varchar(255) not null,
    contact_mobile varchar(255) not null,
    updated_at     datetime     not null,
    primary key (id),
    constraint teams_name_uindex
        unique (name)
);

