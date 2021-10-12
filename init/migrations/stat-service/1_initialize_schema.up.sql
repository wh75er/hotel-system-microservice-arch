CREATE TABLE IF NOT EXISTS stats(
    id serial PRIMARY KEY,
    RoomsAmount INT NOT NULL,
    ReservationsAmount INT NOT NULL
);

INSERT INTO stats(RoomsAmount, ReservationsAmount) VALUES(0, 0);
