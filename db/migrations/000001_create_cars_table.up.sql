CREATE TYPE car_category AS ENUM (
    'Truck', 
    'Sedan', 
    'SUV');

CREATE TABLE IF NOT EXISTS cars(
    id VARCHAR(50) PRIMARY KEY,
    make VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    package VARCHAR(50) NOT NULL,
    year INT NOT NULL,
    category car_category NOT NULL,  
    mileage INT NOT NULL,
    price_cents INT NOT NULL   
);