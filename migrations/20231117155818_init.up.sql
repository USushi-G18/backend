
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

CREATE TABLE plate (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price DECIMAL NOT NULL,
    category_id INT NOT NULL REFERENCES category(id),
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
    allergen_id INT REFERENCES allergen(id)
);

CREATE TABLE plate_ingredient (
    plate_id INT REFERENCES plate(id),
    ingredient_id INT REFERENCES ingredient(id),
    PRIMARY KEY (plate_id, ingredient_id)
);

-- Kitchen

CREATE TYPE command_status as ENUM ('Ordered', 'Preparing', 'Prepared', 'Delivered');

CREATE TABLE session (
    id SERIAL PRIMARY KEY,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP,
    table_number INT NOT NULL,
    menu menu NOT NULL,
    seating INT NOT NULL
);

CREATE TABLE command (
    session_id INT NOT NULL REFERENCES session(id),
    plate_id INT NOT NULL REFERENCES plate(id),
    at TIMESTAMP NOT NULL,
    quantity INT NOT NULL,
    status command_status NOT NULL,
    PRIMARY KEY (session_id, plate_id, at)
);

-- Login

CREATE TYPE sushi_user_type as ENUM ('Employee', 'Admin', 'Client');

CREATE TABLE sushi_user (
    user_type sushi_user_type PRIMARY KEY,
    password TEXT NOT NULL 
);


-- u-sushi hash
-- $argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4

insert into sushi_user (user_type, password) values 
('Admin', '$argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4'),
('Client', '$argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4'),
('Employee', '$argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4');