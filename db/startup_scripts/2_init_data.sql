
--SET NAMES 'utf8';

INSERT INTO onboardings (title, image_url) VALUES ('Удобный поиск и фильтрация', 'https://yaroslavs-imac.local:80/static/onboarding/1.jpg'), ('Сохраняйте рецепты в избранное, создавайте новые и делитесь', 'https://yaroslavs-imac.local:80/static/onboarding/2.jpg'), ('Сканируйте холодильник и проверяйте наличие и срок годности продуктов', 'https://yaroslavs-imac.local:80/static/onboarding/3.jpg');
INSERT INTO users (id, name, image_url) VALUES ('WPLfq0KRbMg7ziylPtR3d2Wditr1', 'Дмитрий Возников', 'https://yaroslavs-imac.local:80/static/avatars/1.jpg');

INSERT INTO recipes (id, title, cooktime, calories, rating, title_image_url, author_id) VALUES (1, 'Новогодний очень вкусный салат', 235, 100, 5, 'https://yaroslavs-imac.local:80/static/recipes/салат.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1'), (2, 'Тыквенный суп котик мяу мяу', 55, 110, 4, 'https://yaroslavs-imac.local:80/static/recipes/тыкв_суп.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1');

INSERT INTO recipe_compilations VALUES (1, 'Аюрведические блюда');
INSERT INTO recipes (id, title, cooktime, calories, rating, title_image_url, author_id) VALUES (3, 'Кичари', 70, 50, 5, 'https://yaroslavs-imac.local:80/static/recipes/кичари.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1'), (4, 'Чечевичные котлеты', 60, 90, 3, 'https://yaroslavs-imac.local:80/static/recipes/чечевич_котлеты.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1');
INSERT INTO recipe_compilations_recipes VALUES (1, 3, 1), (2, 4, 1);

INSERT INTO recipes (id, title, cooktime, calories, rating, title_image_url, author_id) VALUES (5, 'Перловый суп с морским вкусом', 160, 80, 5, 'https://yaroslavs-imac.local:80/static/recipes/перл_суп.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1');
INSERT INTO promo_recipes VALUES (1, 5);

INSERT INTO units VALUES (1, 'г'), (2, 'мл'), (3, 'шт');

INSERT INTO ingredients VALUES (1, 'Картофель'), (2, 'Морковь'), (3, 'Вода'), (4, 'Рис'), (5, 'Имбирь');

INSERT INTO recipes_ingredients VALUES (1, 2, 1, 1, 400), (2, 2, 2, 1, 300), (3, 2, 3, 2, 1500);
INSERT INTO recipes_ingredients VALUES (4, 1, 1, 3, 2);
INSERT INTO recipes_ingredients VALUES (5, 5, 1, 1, 400), (6, 5, 3, 2, 2000), (7, 5, 4, 1, 350), (8, 5, 5, 3, 2);

INSERT INTO recipe_medias VALUES (2, 'https://static.sobaka.ru/images/image/01/63/21/50/_normal.jpeg?v=1671365310', 1), (2, 'https://ir.ozone.ru/s3/multimedia-m/c1000/6275130526.jpg', 1);
INSERT INTO recipe_medias VALUES (1, 'http://jplayer.org/video/m4v/Finding_Nemo_Teaser.m4v', 2);
INSERT INTO recipe_medias VALUES (5, 'https://www.purina.ru/sites/default/files/styles/nppe_breed_selector_500/public/2020-04/french_bulldog.jpg?itok=CWfOzyGk', 1), (5, 'http://jplayer.org/video/m4v/Incredibles_Teaser_320x144_h264aac.m4v', 2), (5, 'https://www.bethowen.ru/upload/medialibrary/ecc/eccb9676c5dbd1524849166d27627e24.jpg', 1);

INSERT INTO recipes_steps VALUES (1, 1, 'Нарежь картоху для парилки', 2), (2, 2, 'Закинь вариться, пока не станет мягкой', 2), (3, 3, 'Произвольно намешай еще чего-нибудь, чтобы было вкусно', 2);
INSERT INTO recipes_steps VALUES (4, 1, 'Рецепт с одним пунктом, чтобы был', 4);
INSERT INTO recipes_steps VALUES (5, 1, 'Мяу', 5), (6, 2, 'Длинный-длинный текст, чтобы проверить многострочное отображение пункта приготовления пищи для поедания и удовлетворения примитивных потребностей организма, продлевающих бытие в этом чудесном мире', 5), (7, 3, 'Мяу-мяу', 5);
