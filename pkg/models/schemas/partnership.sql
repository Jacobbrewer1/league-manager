create table partnership
(
    id          int not null auto_increment,
    player_a_id int not null,
    player_b_id int not null,
    team_id     int not null,
    primary key (id),
    constraint partnership_player_a_id_player_b_id_uindex
        unique (player_a_id, player_b_id, team_id),
    constraint partnership_player_id_fk
        foreign key (player_b_id) references player (id),
    constraint partnership_player_id_fk2
        foreign key (player_a_id) references player (id),
    constraint partnership_team_id_fk
        foreign key (team_id) references team (id)
);

