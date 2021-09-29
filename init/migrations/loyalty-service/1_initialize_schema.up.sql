CREATE TABLE IF NOT EXISTS loyalty(
    id serial PRIMARY KEY,
    userUuid UUID UNIQUE NOT NULL,
    status VARCHAR(250),
    discount INTEGER NOT NULL,
    contributionAmount INTEGER NOT NULL,
);
