CREATE TABLE Company
(
    id                uuid primary key,
    name              varchar(255) not null,
    max_ask           bigint       not null,
    min_ask           bigint       not null,
    max_ask_different bigint       not null,
    max_bid           bigint       not null,
    min_bid           bigint       not null,
    max_bid_different bigint       not null,
    max_ask_bid_diff  bigint       not null
);
