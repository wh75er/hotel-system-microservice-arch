CREATE TABLE IF NOT EXISTS hotels(
    id serial PRIMARY KEY,
    hotelUuid UUID UNIQUE NOT NULL,
    photos UUID[],
    name VARCHAR(250) NOT NULL,
    description VARCHAR(1000) NOT NULL,
    creationDate TIMESTAMP NOT NULL,
    country VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    address VARCHAR(250) NOT NULL,
    isReady BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms(
    id serial PRIMARY KEY,
    roomUuid UUID UNIQUE NOT NULL,
    hotelUuid UUID NOT NULL REFERENCES hotels(hotelUuid),
    RoomType VARCHAR(250) NOT NULL,
    Amount INTEGER NOT NULL,
    Beds INTEGER NOT NULL,
    offers text[],
    nightPrice REAL NOT NULL,
    creationDate TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS reviews(
    id serial PRIMARY KEY,
    reviewUuid UUID UNIQUE NOT NULL,
    hotelUuid UUID NOT NULL REFERENCES hotels(hotelUuid),
    text VARCHAR(1500) NOT NULL,
    userUuid UUID NOT NULL,
    isAnonymous BOOLEAN NOT NULL,
    photos UUID[],
    creationDate TIMESTAMP NOT NULL
);
