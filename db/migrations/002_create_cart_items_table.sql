-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cart_items (
    id 		 BIGSERIAL PRIMARY KEY,
    cart_id  BIGINT,
    product  VARCHAR(40),
    quantity INT,

    FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart_items;
-- +goose StatementEnd
