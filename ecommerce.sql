CREATE TABLE IF NOT EXISTS registered_user {
  id SERIAL PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  last_name VARCHAR(45) NOT NULL,
  e_mail VARCHAR(100) NOT NULL,
  country VARCHAR(45) NOT NULL,
  direction VARCHAR(200) NOT NULL,
  zip_code VARCHAR(30) NOT NULL,
  phone VARCHAR(45) NOT NULL,
  is_admin BOOLEAN NOT NULL,
  password VARCHAR(30) NOT NULL
};

CREATE TABLE IF NOT EXISTS payment_method_type {
  id SERIAL PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  UNIQUE(name)
};

CREATE TABLE IF NOT EXISTS payment_method {
  id SERIAL PRIMARY KEY,
  account_info VARCHAR(100) NOT NULL,
  type INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  FOREIGN KEY(user_id) REFERENCES registered_user(id),  
  FOREIGN KEY(type) REFERENCES payment_method_type(id)
};

CREATE TABLE IF NOT EXISTS purchase {
  id SERIAL PRIMARY KEY,
  direction_to_send VARCHAR(200) NOT NULL,
  zip_code_to_send VARCHAR(30) NOT NULL,
  buyer INTEGER NOT NULL,
  payment_method_id INTEGER NOT NULL,
  UNIQUE(payment_method),
  FOREIGN KEY (buyer) REFERENCES registered_user(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
};

CREATE TABLE IF NOT EXISTS product_category {
  id SERIAL PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  UNIQUE(name)
};

CREATE TABLE IF NOT EXISTS product_offer {
  id SERIAL PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  description VARCHAR(200) NOT NULL,
  price REAL NOT NULL
};

CREATE TABLE IF NOT EXISTS purchase_details {
  product_offer_id INTEGER NOT NULL,
  purchase_id INTEGER NOT NULL,
  FOREIGN KEY (product_offer_id) REFERENCES product_offer(id),
  FOREIGN KEY (purchase_id) REFERENCES purchase(id),
  PRIMARY KEY(product_offer, purchase)
};