create table score
(
    id               int not null auto_increment,
    game_id         int not null,
    partnership_id   int not null,
    first_set_score  int not null,
    second_set_score int not null,
    third_set_score  int null,
    primary key (id),
    constraint score_matches_id_fk
        foreign key (game_id) references game (id),
    constraint score_partnership_id_fk
        foreign key (partnership_id) references partnership (id)
);

