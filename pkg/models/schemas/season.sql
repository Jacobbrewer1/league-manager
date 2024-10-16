create table season
(
    id   int         not null auto_increment,
    name varchar(25) not null,
    primary key (id),
    constraint season_name_uindex
        unique (name)
);

