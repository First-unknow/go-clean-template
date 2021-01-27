
DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id INT PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR (255)
);

INSERT INTO users (id, email, first_name, last_name) VALUES ( 1020001, 'test@test.test', 'test', 'test');
