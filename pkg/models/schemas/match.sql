create table `match`
(
    id               int      not null auto_increment,
    season_id        int      not null,
    home_partners_id int      not null,
    away_partners_id int      not null,
    match_date       datetime not null,
    primary key (id),
    constraint match_season_id_fk
        foreign key (season_id) references season (id),
    constraint matches_partnership_id_fk
        foreign key (home_partners_id) references partnership (id),
    constraint matches_partnership_id_fk2
        foreign key (away_partners_id) references partnership (id)
);

