CREATE TABLE IF NOT EXISTS reservations(
    id serial PRIMARY KEY,
    ReservationUuid UUID UNIQUE NOT NULL,
    RoomUuid UUID NOT NULL,
    UserUuid UUID NOT NULL,
    PaymentUuid UUID,
    Date DATE NOT NULL,
    Status VARCHAR(64) NOT NULL
);
