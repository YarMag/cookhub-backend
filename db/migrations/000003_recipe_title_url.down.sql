ALTER TABLE IF EXISTS cookhubdb.recipe DROP COLUMN title_image_url;
DROP TABLE IF EXISTS cookhubdb.user;

CREATE TABLE IF NOT EXISTS cookhubdb.author
(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	image_url VARCHAR(300) NOT NULL
);

ALTER TABLE IF EXISTS cookhubdb.recipe DROP COLUMN author_id;
ALTER TABLE IF EXISTS cookhubdb.recipe ADD COLUMN author_id SERIAL NOT NULL;
ALTER TABLE IF EXISTS cookhubdb.recipe ADD CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES cookhubdb.author(id)

DROP TABLE IF EXISTS cookhubdb.favorite_recipes;