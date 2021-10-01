CREATE TABLE IF NOT EXISTS Users(
    id serial PRIMARY KEY,
    userUuid UUID UNIQUE NOT NULL,
    login VARCHAR(24) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL,
    role VARCHAR(100) NOT NULL
);
