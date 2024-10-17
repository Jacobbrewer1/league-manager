create table game
(
    id               int      not null auto_increment,
    season_id        int      not null,
    home_partners_id int      not null,
    away_partners_id int      not null,
    match_date       datetime not null,
    winning_team     enum ('HOME', 'AWAY') not null,
    primary key (id),
    constraint game_season_id_fk
        foreign key (season_id) references season (id),
    constraint game_partnership_id_fk
        foreign key (home_partners_id) references partnership (id),
    constraint game_partnership_id_fk2
        foreign key (away_partners_id) references partnership (id)
);

