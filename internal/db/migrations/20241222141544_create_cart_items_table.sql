-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INT,
    product VARCHAR(40),
    quantity INT,

    FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart_items;
-- +goose StatementEnd
