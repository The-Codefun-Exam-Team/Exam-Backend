-- +goose Up

ALTER TABLE debug_problems
ADD mindiff INT NOT NULL DEFAULT 0;

-- +goose Down

ALTER TABLE debug_problems
DROP COLUMN mindiff;
