ALTER TABLE games ADD CONSTRAINT games_price_check CHECK (price > 0);
ALTER TABLE games ADD CONSTRAINT title_color_check CHECK (title is not null);
ALTER TABLE games ADD CONSTRAINT ages_length_check CHECK (ages is not null);