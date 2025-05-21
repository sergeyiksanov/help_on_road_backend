-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE users ADD auto_model VARCHAR(100);
ALTER TABLE users ADD auto_gos_number VARCHAR(100);
ALTER TABLE users ADD vin_code VARCHAR(100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE users DROP COLUMN auto_model;
ALTER TABLE users DROP COLUMN auto_gos_number;
ALTER TABLE users DROP COLUMN vin_code;
-- +goose StatementEnd
