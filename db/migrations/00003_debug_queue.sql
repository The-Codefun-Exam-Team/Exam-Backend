-- +goose Up

CREATE TABLE debug_queue (
    rid INT NOT NULL,
    drid INT NOT NULL,

    PRIMARY KEY(rid),
    UNIQUE(drid)
);

-- +goose Down

DROP TABLE IF EXISTS debug_queue;