ALTER TABLE IF EXISTS cookhubdb.recipe ADD COLUMN title_image_url VARCHAR(100);
ALTER TABLE IF EXISTS cookhubdb.recipe DROP COLUMN author_id;

DROP TABLE IF EXISTS cookhubdb.author;
CREATE TABLE IF NOT EXISTS cookhubdb.user
(
	id VARCHAR(128) PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	image_url VARCHAR(300) NOT NULL
);

ALTER TABLE IF EXISTS cookhubdb.recipe ADD COLUMN author_id VARCHAR(128) NOT NULL;
ALTER TABLE IF EXISTS cookhubdb.recipe ADD CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES cookhubdb.user(id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS cookhubdb.favorite_recipes
(
	user_id VARCHAR(128) REFERENCES cookhubdb.user(id) NOT NULL,
	recipe_id SERIAL REFERENCES cookhubdb.recipe(id) NOT NULL
);

CREATE INDEX ON cookhubdb.favorite_recipes (user_id, recipe_id);