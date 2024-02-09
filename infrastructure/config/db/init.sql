CREATE DATABASE IF NOT EXISTS history;

USE history;

CREATE TABLE IF NOT EXISTS arith_history
(
    date      DATE         NOT NULL,
    answer    INT          NOT NULL,
    operation VARCHAR(255) NOT NULL
);


CREATE DATABASE IF NOT EXISTS test_history;

USE test_history;

CREATE TABLE arith_history
(
    date      DATE         NOT NULL,
    answer    INT          NOT NULL,
    operation VARCHAR(255) NOT NULL
);
