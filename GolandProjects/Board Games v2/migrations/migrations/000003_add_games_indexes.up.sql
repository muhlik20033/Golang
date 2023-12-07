CREATE INDEX IF NOT EXISTS games_title_idx ON games USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS games_color_idx ON games USING GIN (to_tsvector('simple', color));
