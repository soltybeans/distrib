CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE players (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    player_name VARCHAR(64) NOT NULL
);

CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    move_number INTEGER NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP,
    last_move_date TIMESTAMP,
    player_one UUID NOT NULL REFERENCES players(id),
    player_two UUID NOT NULL REFERENCES players(id)
);

CREATE INDEX recent_moves on games(last_move_date DESC, id) INCLUDE (move_number, player_one, player_two, created_date, end_date)