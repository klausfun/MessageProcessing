CREATE TABLE status
(
    id   serial      not null unique,
    name varchar(20) not null unique
);

INSERT INTO status (name)
VALUES ('current'),
       ('completed');

CREATE TABLE message
(
    id        serial                     not null unique,
    content   varchar(4239)              not null,
    status_id int references status (id) not null default 1
);
