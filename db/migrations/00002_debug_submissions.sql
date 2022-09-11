-- +goose Up

CREATE TABLE debug_submissions (
    drid INT auto_increment NOT NULL,
    dpid INT NOT NULL,
    tid INT NOT NULL,
    language VARCHAR(8) NOT NULL,
    submittime INT NOT NULL,
    score DOUBLE NOT NULL,
    diff DOUBLE NOT NULL,
    code LONGTEXT NOT NULL,

    PRIMARY KEY(drid),

    FOREIGN KEY(dpid) REFERENCES debug_problems(dpid),

    FOREIGN KEY(tid) REFERENCES teams(tid)
);

-- +goose Down

DROP TABLE IF EXISTS debug_submissions;