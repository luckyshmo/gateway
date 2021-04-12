CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE raw_data
(
    id            uuid DEFAULT uuid_generate_v4 () not null unique,
    time_cr       timestamp not null,
    data_r        bytea not null
);

CREATE TABLE valid_data
(
    id            uuid DEFAULT uuid_generate_v4 () not null unique,
    dev_eui       varchar(255) not null,
    time_cr       timestamp not null,
    time_p        timestamp not null,
    data_f        bytea not null,
    raw_data      bytea not null
);