CREATE TABLE drivers (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(11) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE vehicles (
    id SERIAL PRIMARY KEY,
    license_plate VARCHAR(7) NOT NULL UNIQUE,
    model VARCHAR(50) NOT NULL
);

CREATE TABLE driver_vehicles (
    driver_id INT NOT NULL,
    vehicle_id INT NOT NULL,
    PRIMARY KEY (driver_id, vehicle_id),
    FOREIGN KEY (driver_id) REFERENCES drivers(id) ON DELETE CASCADE,
    FOREIGN KEY (vehicle_id) REFERENCES vehicles(id) ON DELETE CASCADE
);