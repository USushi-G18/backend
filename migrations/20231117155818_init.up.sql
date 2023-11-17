
-- Client 

CREATE TYPE menu as ENUM ('carte', 'dinner', 'lunch');

CREATE TABLE menu_price (
    menu menu PRIMARY KEY,
    price DECIMAL NOT NULL 
);

CREATE TABLE image (
    id SERIAL PRIMARY KEY,
    -- image base64 encoding
    image TEXT NOT NULL 
);

CREATE TABLE category (
    name TEXT PRIMARY KEY
);

CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price DECIMAL NOT NULL,
    category TEXT NOT NULL REFERENCES category(name),
    menu menu NOT NULL,
    description text,
    image_id INT REFERENCES image(id),
    -- order_limit null means no limit
    order_limit INT, 
    -- ex. nighiri 2 piec.
    pieces INT NOT NULL
);


CREATE TABLE allergen (
    name TEXT PRIMARY KEY 
);

CREATE TABLE ingredient (
    name TEXT PRIMARY KEY,
    allergen TEXT REFERENCES allergen(name)
);

CREATE TABLE product_ingredient (
    product_id INT REFERENCES product(id),
    ingredient_name TEXT REFERENCES ingredient(name),
    PRIMARY KEY (product_id, ingredient_name)
);

-- Kitchen

CREATE TYPE command_status as ENUM ('ordered', 'preparing', 'prepared', 'delivered');

CREATE TABLE session (
    id SERIAL PRIMARY KEY,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP,
    table_number INT NOT NULL
);

CREATE TABLE command (
    session INT NOT NULL REFERENCES session(id),
    product_id INT NOT NULL REFERENCES product(id),
    at TIMESTAMP NOT NULL,
    quantity INT NOT NULL,
    status command_status NOT NULL
);

-- Login

CREATE TYPE sushi_user_type as ENUM ('kitchen', 'admin', 'client');

CREATE TABLE sushi_user (
    user_type sushi_user_type PRIMARY KEY,
    password TEXT NOT NULL 
);