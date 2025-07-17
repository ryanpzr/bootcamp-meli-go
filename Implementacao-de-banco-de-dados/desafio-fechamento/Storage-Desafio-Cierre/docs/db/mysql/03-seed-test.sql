USE `fantasy_products_test`;

INSERT INTO customers(id, last_name, first_name, `condition`) VALUES (1,  'Pereira', 'Ryan',1);
INSERT INTO customers(id, last_name, first_name, `condition`) VALUES (2,  'Frota', 'Marcos',0);

INSERT INTO products(id, description, price) VALUES (1, 'Truffle Cups - Red', 20.55);
INSERT INTO products(id, description, price) VALUES (2, 'Sword Pick Asst', 33.99);

INSERT INTO invoices(id, datetime, customer_id, total) VALUES (1, '2004-10-07', 1, 23.0);
INSERT INTO invoices(id, datetime, customer_id, total) VALUES (2, '2008-11-24', 2, 34.5);

INSERT INTO sales(id, quantity, invoice_id, product_id) VALUES (1, 4, 2, 1);
INSERT INTO sales(id, quantity, invoice_id, product_id) VALUES (2, 12, 2, 1);
INSERT INTO sales(id, quantity, invoice_id, product_id) VALUES (3, 8, 1, 2);
INSERT INTO sales(id, quantity, invoice_id, product_id) VALUES (4, 12, 2, 1);