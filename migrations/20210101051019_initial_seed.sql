-- +goose Up
-- +goose StatementBegin
CREATE TABLE ping_timestamp (
    id SERIAL,
    occurred TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ping_timestamp;
-- +goose StatementEnd
