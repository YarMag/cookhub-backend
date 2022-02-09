CREATE TABLE IF NOT EXISTS cookhubdb.onboardings
(
	onboarding_id SERIAL PRIMARY KEY,
	title VARCHAR (100) NOT NULL,
	image_url VARCHAR(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS cookhubdb.users
(
	id VARCHAR(128) PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	image_url VARCHAR(300)
);

CREATE TABLE IF NOT EXISTS cookhubdb.recipes
(
	id SERIAL PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	cooktime INT NOT NULL,
	calories DECIMAL NOT NULL,
	rating DECIMAL(2,1),
	author_id VARCHAR(128) NOT NULL REFERENCES cookhubdb.users(id) ON DELETE CASCADE,
	CONSTRAINT CHK_rating CHECK (rating >= 0 AND rating <= 5)
);

CREATE TABLE IF NOT EXISTS cookhubdb.recipes_steps
(
	id SERIAL PRIMARY KEY,
	step INT NOT NULL,
	description VARCHAR(500) NOT NULL,
	recipe_id SERIAL NOT NULL REFERENCES cookhubdb.recipes(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cookhubdb.favorite_recipes
(
	user_id VARCHAR(128) NOT NULL REFERENCES cookhubdb.users(id) ON DELETE CASCADE,
	recipe_id SERIAL NOT NULL REFERENCES cookhubdb.recipes(id) ON DELETE CASCADE
);

CREATE INDEX ON cookhubdb.favorite_recipes (user_id, recipe_id);