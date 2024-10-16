create table `match`
(
    id               int      not null auto_increment,
    home_partners_id int      not null,
    away_partners_id int      not null,
    match_date       datetime not null,
    primary key (id),
    constraint matches_partnership_id_fk
        foreign key (home_partners_id) references partnership (id),
    constraint matches_partnership_id_fk2
        foreign key (away_partners_id) references partnership (id)
);

