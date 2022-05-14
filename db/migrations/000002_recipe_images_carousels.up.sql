ALTER TABLE cookhubdb.recipes ADD COLUMN IF NOT EXISTS title_image_url VARCHAR(100) DEFAULT NULL;

CREATE TABLE IF NOT EXISTS cookhubdb.recipe_compilations 
(
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS cookhubdb.recipe_compilations_recipes
(
	id SERIAL PRIMARY KEY,
	id_recipe SERIAL NOT NULL REFERENCES cookhubdb.recipes(id) ON DELETE CASCADE,
	id_compilation SERIAL NOT NULL REFERENCES cookhubdb.recipe_compilations(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cookhubdb.promo_recipes
(
	id SERIAL PRIMARY KEY,
	id_recipe SERIAL NOT NULL REFERENCES cookhubdb.recipes(id) ON DELETE CASCADE
);