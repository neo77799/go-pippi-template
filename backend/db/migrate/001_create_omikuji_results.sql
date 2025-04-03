-- +goose Up
CREATE TABLE omikuji_results (
    id SERIAL PRIMARY KEY,
    result TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS omikuji_results;
