DROP Table IF EXISTS all_ports;

CREATE TABLE IF NOT EXISTS all_ports(
    id SERIAL NOT NULL PRIMARY KEY,
    name        VARCHAR,
    ref_name    VARCHAR,
    city        VARCHAR,
    country     VARCHAR,
    alias       VARCHAR,
    regions     VARCHAR,
    coordinates VARCHAR,
    province    VARCHAR,
    timezone    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    unlocs      VARCHAR,
    code        VARCHAR
);