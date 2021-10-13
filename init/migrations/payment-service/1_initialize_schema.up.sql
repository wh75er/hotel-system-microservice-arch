CREATE TABLE IF NOT EXISTS payments(
    id serial PRIMARY KEY,
    paymentUuid UUID UNIQUE NOT NULL,
    userUuid UUID NOT NULL,
    status VARCHAR(250) NOT NULL,
    price REAL NOT NULL,
    timeUpdated DATE NOT NULL
);
