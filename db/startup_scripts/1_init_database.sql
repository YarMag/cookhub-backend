GRANT ALL PRIVILEGES ON  DATABASE cookhubdb TO ymagin;

--ALTER DATABASE cookhubdb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

--CREATE SCHEMA cookhubdb;
--ALTER USER ymagin WITH PASSWORD 'magin13';

CREATE TABLE onboardings
(
	onboarding_id SERIAL PRIMARY KEY,
	title VARCHAR (100) NOT NULL,
	image_url VARCHAR(300) NOT NULL
);

CREATE TABLE users
(
	id VARCHAR(128) PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	image_url VARCHAR(300)
);

CREATE TABLE recipes
(
	id SERIAL PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	cooktime INT NOT NULL,
	calories DECIMAL NOT NULL,
	rating DECIMAL(2,1),
	title_image_url VARCHAR(100) DEFAULT NULL,
	author_id VARCHAR(128) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	CONSTRAINT CHK_rating CHECK (rating >= 0 AND rating <= 5)
);

CREATE TABLE recipes_steps
(
	id SERIAL PRIMARY KEY,
	step INT NOT NULL,
	description VARCHAR(500) NOT NULL,
	recipe_id SERIAL NOT NULL REFERENCES recipes(id) ON DELETE CASCADE
);

CREATE TABLE favorite_recipes
(
	user_id VARCHAR(128) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	recipe_id SERIAL NOT NULL REFERENCES recipes(id) ON DELETE CASCADE
);

CREATE INDEX ON favorite_recipes (user_id, recipe_id);

CREATE TABLE recipe_compilations 
(
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL
);

CREATE TABLE recipe_compilations_recipes
(
	id SERIAL PRIMARY KEY,
	id_recipe SERIAL NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
	id_compilation SERIAL NOT NULL REFERENCES recipe_compilations(id) ON DELETE CASCADE
);

CREATE TABLE promo_recipes
(
	id SERIAL PRIMARY KEY,
	id_recipe SERIAL NOT NULL REFERENCES recipes(id) ON DELETE CASCADE
);