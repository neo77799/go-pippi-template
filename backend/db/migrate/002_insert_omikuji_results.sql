-- +goose Up
INSERT INTO omikuji_results (result) VALUES
('大吉'),
('吉'),
('中吉'),
('小吉'),
('凶'),
('大凶');

-- +goose Down
DELETE FROM omikuji_results WHERE result IN (
    '大吉',
    '吉',
    '中吉',
    '小吉',
    '凶',
    '大凶'
);
