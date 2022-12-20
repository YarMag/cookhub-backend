
--SET NAMES 'utf8';

INSERT INTO onboardings (title, image_url) VALUES ('Удобный поиск и фильтрация', 'https://yaroslavs-imac.local:80/static/onboarding/1.jpg'), ('Сохраняйте рецепты в избранное, создавайте новые и делитесь', 'https://yaroslavs-imac.local:80/static/onboarding/2.jpg'), ('Сканируйте холодильник и проверяйте наличие и срок годности продуктов', 'https://yaroslavs-imac.local:80/static/onboarding/3.jpg');
INSERT INTO users (id, name, image_url) VALUES ('WPLfq0KRbMg7ziylPtR3d2Wditr1', 'Дмитрий Возников', 'https://yaroslavs-imac.local:80/static/avatars/1.jpg');

INSERT INTO recipes (id, title, cooktime, calories, rating, title_image_url, author_id) VALUES (1, 'Новогодний очень вкусный салат', 235, 100, 5, 'https://yaroslavs-imac.local:80/static/recipes/салат.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1'), (2, 'Тыквенный суп котик мяу мяу', 55, 110, 4, 'https://yaroslavs-imac.local:80/static/recipes/тыкв_суп.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1');

INSERT INTO recipe_compilations VALUES (1, 'Аюрведические блюда');
INSERT INTO recipes (id, title, cooktime, calories, rating, title_image_url, author_id) VALUES (3, 'Кичари', 70, 50, 5, 'https://yaroslavs-imac.local:80/static/recipes/кичари.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1'), (4, 'Чечевичные котлеты', 60, 90, 3, 'https://yaroslavs-imac.local:80/static/recipes/чечевич_котлеты.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1');
INSERT INTO recipe_compilations_recipes VALUES (1, 3, 1), (2, 4, 1);

INSERT INTO recipes (id, title, cooktime, calories, rating, title_image_url, author_id) VALUES (5, 'Перловый суп с морским вкусом', 160, 80, 5, 'https://yaroslavs-imac.local:80/static/recipes/перл_суп.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1');
INSERT INTO promo_recipes VALUES (1, 5);
