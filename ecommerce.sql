DROP TABLE IF EXISTS registered_user CASCADE;
CREATE TABLE IF NOT EXISTS registered_user (
    id SERIAL PRIMARY KEY,
    photo_url VARCHAR(100) NOT NULL,
    name VARCHAR(45) NOT NULL,
    last_name VARCHAR(45) NOT NULL,
    e_mail VARCHAR(100) NOT NULL,
    country VARCHAR(45) NOT NULL,
    direction VARCHAR(200) NOT NULL,
    zip_code VARCHAR(30) NOT NULL,
    phone VARCHAR(45) NOT NULL,
    is_admin BOOLEAN NOT NULL,
    password VARCHAR(30) NOT NULL,
    UNIQUE(e_mail)
);
INSERT INTO registered_user (photo_url, name, 
last_name, e_mail, country, direction, zip_code,
phone, is_admin, password) VALUES 
('url1', 'Manuel A.', 'Mena S.', 'mena.0837@gmail.com', 'Colombia', 'Poblado', '5465', '5555', TRUE, 'admin'),
('url2', 'Olga', 'Lucila', 'olga.lucilas@yahoo.es', 'Colombia', 'Poblado', '5465', '5555', TRUE, 'admin'),
('url3', 'Marcela', 'Mena', 'marcela.mena@correounivalle.edu.co', 'Colombia', 'Manzanares', '5432', '3134567894', FALSE, 'diana')
;

DROP TABLE IF EXISTS payment_method_type CASCADE;
CREATE TABLE IF NOT EXISTS payment_method_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    UNIQUE(name)
);
INSERT INTO payment_method_type (name) VALUES 
('Tarjeta debito'), ('Paypal'), ('Tarjeta credito')
;

DROP TABLE IF EXISTS payment_method CASCADE;
CREATE TABLE IF NOT EXISTS payment_method (
    id SERIAL PRIMARY KEY,
    account_info VARCHAR(100) NOT NULL,
    type INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES registered_user(id),  
    FOREIGN KEY(type) REFERENCES payment_method_type(id)
);
INSERT INTO payment_method (account_info, type, user_id) VALUES 
('12345678',1,1),
('455-567',2,1),
('456-123',1,2),
('456-789',2,2)
;

DROP TABLE IF EXISTS purchase CASCADE;
CREATE TABLE IF NOT EXISTS purchase (
    id SERIAL PRIMARY KEY,
    direction_to_send VARCHAR(200) NOT NULL,
    zip_code_to_send VARCHAR(30) NOT NULL,
    buyer INTEGER NOT NULL,
    payment_method_id INTEGER NOT NULL,
    UNIQUE(payment_method_id),
    FOREIGN KEY (buyer) REFERENCES registered_user(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
);
INSERT INTO purchase (direction_to_send, zip_code_to_send, buyer, payment_method_id)
VALUES 
('Manzana 22D Casa 53', '53555', 1, 3),
('Cra 14 #14-55', '54678', 2, 4)
;

DROP TABLE IF EXISTS product_category CASCADE;
CREATE TABLE IF NOT EXISTS product_category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    UNIQUE(name)
);
INSERT INTO product_category (name) VALUES
('Electrodomesticos'),
('Automoviles'),
('Casas')
;

DROP TABLE IF EXISTS product_offer CASCADE;
CREATE TABLE IF NOT EXISTS product_offer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    description VARCHAR(200) NOT NULL,
    price REAL NOT NULL,
    principal_image_url VARCHAR(100) NOT NULL
);
INSERT INTO product_offer(name, description, price, principal_image_url) VALUES
('Producto 1', 'Descripcion 1', 54567, 'product1'),
('Producto 2', 'Descripcion 2', 67890, 'product2')
;

DROP TABLE IF EXISTS product_offer_img_urls CASCADE;
CREATE TABLE IF NOT EXISTS product_offer_img_urls (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    img_url VARCHAR(100) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product_offer(id)
);
INSERT INTO product_offer_img_urls(product_id, img_url) VALUES
(1, 'product1-2'),
(1, 'product1-3')
;

DROP TABLE IF EXISTS purchase_details CASCADE;
CREATE TABLE IF NOT EXISTS purchase_details (
    product_offer_id INTEGER NOT NULL,
    purchase_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    FOREIGN KEY (product_offer_id) REFERENCES product_offer(id),
    FOREIGN KEY (purchase_id) REFERENCES purchase(id),
    PRIMARY KEY (product_offer_id, purchase_id)
);
INSERT INTO purchase_details (product_offer_id, purchase_id, quantity) VALUES
(1, 1, 3),
(2, 2, 5)
;