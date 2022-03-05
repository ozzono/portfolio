DROP Table IF EXISTS json_control;

CREATE Table if not exists json_control(
    version SERIAL NOT NULL PRIMARY KEY,
    parsed boolean
)