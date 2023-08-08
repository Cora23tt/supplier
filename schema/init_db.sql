DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS shipments;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    id SERIAL NOT NULL UNIQUE,
    username VARCHAR(225) NOT NULL UNIQUE,
    password_hash VARCHAR(225) NOT NULL,
    first_name VARCHAR(225) NOT NULL UNIQUE,
    last_name VARCHAR(225) NOT NULL,
    email VARCHAR(254) NOT NULL UNIQUE,
    address VARCHAR(225) NOT NULL UNIQUE,
    telephone INT NOT NULL
);



CREATE TABLE categories
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL
);
CREATE TABLE products
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(200) NOT NULL,
    description VARCHAR(225),
    category_id INT REFERENCES categories(id) ON DELETE CASCADE NOT NULL,
    price NUMERIC(15, 2) NOT NULL
);



CREATE TABLE carts
(
  id SERIAL NOT NULL UNIQUE,
  user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);
CREATE TABLE cart_items
(
  cart_id INT REFERENCES carts(id) ON DELETE CASCADE NOT NULL,
  product_id INT REFERENCES products(id) ON DELETE CASCADE NOT NULL
);



CREATE TABLE payments
(
    id SERIAL NOT NULL UNIQUE,
    date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    method VARCHAR(100),
    amount NUMERIC(15, 2) NOT NULL,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);



CREATE TABLE shipments
(
    id SERIAL NOT NULL UNIQUE,
    address VARCHAR(100),
    date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    status VARCHAR(100)
);



CREATE TABLE orders
(
    id SERIAL NOT NULL UNIQUE,
    order_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    total_price NUMERIC(16, 2),
    payment_id INT REFERENCES payments(id) ON DELETE CASCADE NOT NULL,
    shipment_id INT REFERENCES shipments(id) ON DELETE CASCADE NOT NULL
);
CREATE TABLE order_items
(
    order_id INT REFERENCES orders(id) ON DELETE CASCADE NOT NULL,
    product_id INT REFERENCES products(id) ON DELETE CASCADE NOT NULL
);