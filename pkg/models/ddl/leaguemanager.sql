create table player
(
    id         int auto_increment
        primary key,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    email      varchar(255) not null,
    dob        date         not null,
    updated_at datetime     not null
)
    with system versioning;

create table season
(
    id   int auto_increment
        primary key,
    name varchar(25) not null,
    constraint season_name_uindex
        unique (name)
);

create table team
(
    id             int auto_increment
        primary key,
    name           varchar(25)  not null,
    contact_email  varchar(255) not null,
    contact_mobile varchar(255) not null,
    updated_at     datetime     not null,
    constraint teams_name_uindex
        unique (name)
);

create table partnership
(
    id          int auto_increment
        primary key,
    player_a_id int not null,
    player_b_id int not null,
    team_id     int not null,
    constraint partnership_player_a_id_player_b_id_uindex
        unique (player_a_id, player_b_id, team_id),
    constraint partnership_player_id_fk
        foreign key (player_b_id) references player (id),
    constraint partnership_player_id_fk2
        foreign key (player_a_id) references player (id),
    constraint partnership_team_id_fk
        foreign key (team_id) references team (id)
);

create table game
(
    id               int auto_increment
        primary key,
    season_id        int      not null,
    home_partners_id int      not null,
    away_partners_id int      not null,
    match_date       datetime not null,
    winning_team     enum ('HOME', 'AWAY') not null,
    constraint game_season_id_fk
        foreign key (season_id) references season (id),
    constraint games_partnership_id_fk
        foreign key (home_partners_id) references partnership (id),
    constraint games_partnership_id_fk2
        foreign key (away_partners_id) references partnership (id)
);

create table score
(
    id               int auto_increment
        primary key,
    game_id         int not null,
    partnership_id   int not null,
    first_set_score  int not null,
    second_set_score int not null,
    third_set_score  int null,
    constraint score_games_id_fk
        foreign key (game_id) references game (id),
    constraint score_partnership_id_fk
        foreign key (partnership_id) references partnership (id)
);

