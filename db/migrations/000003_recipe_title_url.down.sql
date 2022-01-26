ALTER TABLE IF EXISTS cookhubdb.recipe DROP COLUMN title_image_url;
ALTER TABLE IF EXISTS cookhubdb.user RENAME TO cookhubdb.author;

ALTER TABLE IF EXISTS cookhubdb.recipe ADD COLUMN author_id SERIAL REFERENCES author(id) NOT NULL;

DROP TABLE IF EXISTS cookhubdb.favorite_recipes;