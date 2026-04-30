-- +goose Up
CREATE TABLE IF NOT EXISTS subscribtions (
    id BIGSERIAL PRIMARY KEY,
    service_name TEXT NOT NULL,
    price_in_ruble INTEGER NOT NULL CHECK (price_in_ruble > 0),
    id_user INTEGER NOT NULL,
    start_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    end_date TIMESTAMPTZ
);

INSERT INTO subscribtions (service_name, price_in_ruble, id_user, start_date, end_date) 
VALUES 
    ('Yandex Cloud', 30, 1, NOW(), NOW()),
    ('Wildberies', 15, 2, NOW(), NOW()),
    ('OZON', 16, 4, NOW(), NOW()),
    ('MTS', 20, 5, NOW(), NOW());

-- +goose Down
DROP TABLE IF EXISTS subscribtions;
