create table voter
(
    id  serial
        primary key,
    candidate_1      varchar,
    candidate_2      varchar,
    candidate_3      varchar,
    has_voted        boolean
);

create table voter_reg
(
    candidate_1      uuid,
    candidate_2      uuid,
    candidate_3       uuid,
    can_verify_1 varchar,
    can_verify_2 varchar,
    can_verify_3 varchar,
    health_card varchar,
    name        varchar,
    id          serial
        primary key
);

create table polls
(
    candidate integer
        primary key,
    votes int
);