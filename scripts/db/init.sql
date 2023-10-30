CREATE TABLE shorten_url (
    id SERIAL PRIMARY KEY,
    code VARCHAR(6) NOT NULL,
    original VARCHAR(1000) NOT NULL,
    shorten_url VARCHAR(25) NOT NULL,
    access_number INT NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);