CREATE TABLE IF NOT EXISTS reservations(
    id serial PRIMARY KEY,
    ReservationUUID UNIQUE NOT NULL,
    RoomUuid UUID NOT NULL,
    UserUuid UUID NOT NULL,
    PaymentUuid UUID UNIQUE,
    Date DATE NOT NULL
    Status VARCHAR(64) NOT NULL,
);
