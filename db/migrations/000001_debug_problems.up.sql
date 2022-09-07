BEGIN;

CREATE TABLE debug_problems (
	dpid INT auto_increment NOT NULL,
	code VARCHAR(20) NOT NULL,
	name TINYTEXT NOT NULL,
	status TINYTEXT DEFAULT "Active",
	solved INT NOT NULL DEFAULT 0,
	total INT NOT NULL DEFAULT 0,
	rid INT NOT NULL,
	pid INT NOT NULL,
	language VARCHAR(8) NOT NULL,
	score DOUBLE NOT NULL DEFAULT 0,
	result VARCHAR(8) NOT NULL,

	PRIMARY KEY(dpid),
	UNIQUE(code),

	FOREIGN KEY(rid) REFERENCES runs(rid)
);

COMMIT;
