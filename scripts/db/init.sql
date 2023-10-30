CREATE TABLE shorten_url (
    id SERIAL PRIMARY KEY,
    code VARCHAR(6) NOT NULL,
    original VARCHAR(1000) NOT NULL,
    access_number INT NOT NULL
);