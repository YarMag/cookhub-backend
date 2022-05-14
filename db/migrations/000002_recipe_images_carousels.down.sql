ALTER TABLE cookhubdb.recipes DROP COLUMN IF EXISTS title_image_url;

DROP TABLE IF EXISTS cookhubdb.recipe_compilations;
DROP TABLE IF EXISTS cookhubdb.recipe_compilations_recipes;
DROP TABLE IF EXISTS cookhubdb.promo_recipes;