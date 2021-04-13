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
    p_type        varchar(255) not null,
    time_cr       timestamp not null,
    raw_data      bytea not null,
    dev_eui       varchar(255),
    app_eui       varchar(255),
    ack           boolean,
    data_f        varchar(255),
    dr            varchar(255),
    fcnt          numeric,
    freq          numeric,
    gateway_id    varchar(255),
    port          numeric,
    rssi          numeric,
    snr           numeric, --double precision,
    time_stamp_   numeric,
    type_         varchar(255)
);