-- +goose Up
CREATE TABLE IF NOT EXISTS subscribtions (
    id BIGSERIAL PRIMARY KEY,
    service_name TEXT NOT NULL,
    price_in_ruble INTEGER NOT NULL CHECK (price_in_ruble > 0),
    id_user INTEGER NOT NULL,
    start_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    end_date TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS subscribtions;
