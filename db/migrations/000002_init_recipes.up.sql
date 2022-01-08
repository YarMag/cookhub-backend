CREATE TABLE IF NOT EXISTS cookhubdb.author
(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	image_url VARCHAR(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS cookhubdb.recipe
(
	id SERIAL PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	cooktime INT NOT NULL,
	calories DECIMAL NOT NULL,
	rating DECIMAL,
	author_id SERIAL REFERENCES author(id) NOT NULL,
	CONSTRAINT CHK_rating CHECK (rating >= 0 AND rating <= 5)
);

CREATE TABLE IF NOT EXISTS cookhubdb.recipe_step
(
	id SERIAL PRIMARY KEY,
	step INT NOT NULL,
	description VARCHAR(500) NOT NULL,
	recipe_id SERIAL REFERENCES recipe(id) NOT NULL
);