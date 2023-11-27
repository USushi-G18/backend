
-- Client 

CREATE TABLE image (
    id SERIAL PRIMARY KEY,
    -- image base64 encoding
    image TEXT NOT NULL 
);

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE 
);

CREATE TYPE menu as ENUM ('Carte', 'Dinner', 'Lunch');

CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price DECIMAL NOT NULL,
    category INT NOT NULL REFERENCES category(id),
    menu menu NOT NULL,
    description text,
    image_id INT REFERENCES image(id),
    -- order_limit null means no limit
    order_limit INT, 
    -- ex. nigiri 2 piec.
    pieces INT NOT NULL
);


CREATE TABLE allergen (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE 
);

CREATE TABLE ingredient (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE,
    allergen INT REFERENCES allergen(id)
);

CREATE TABLE product_ingredient (
    product_id INT REFERENCES product(id),
    ingredient_id INT REFERENCES ingredient(id),
    PRIMARY KEY (product_id, ingredient_id)
);

-- Kitchen

CREATE TYPE command_status as ENUM ('Ordered', 'Preparing', 'Prepared', 'Delivered');

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

CREATE TYPE sushi_user_type as ENUM ('Kitchen', 'Admin', 'Client');

CREATE TABLE sushi_user (
    user_type sushi_user_type PRIMARY KEY,
    password TEXT NOT NULL 
);