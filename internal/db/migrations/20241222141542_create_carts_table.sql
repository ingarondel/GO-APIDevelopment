-- +goose Up
-- +goose StatementBegin
CREATE TABLE carts (
    id BIGSERIAL PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE carts;
-- +goose StatementEnd
