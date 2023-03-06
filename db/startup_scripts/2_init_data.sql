
--SET NAMES 'utf8';

INSERT INTO onboardings (title, image_url) VALUES ('Удобный поиск и фильтрация', 'https://yaroslavs-imac.local:80/static/onboarding/1.jpg'), ('Сохраняйте рецепты в избранное, создавайте новые и делитесь', 'https://yaroslavs-imac.local:80/static/onboarding/2.jpg'), ('Сканируйте холодильник и проверяйте наличие и срок годности продуктов', 'https://yaroslavs-imac.local:80/static/onboarding/3.jpg');
INSERT INTO users (id, name, image_url) VALUES ('WPLfq0KRbMg7ziylPtR3d2Wditr1', 'Дмитрий Возников', 'https://yaroslavs-imac.local:80/static/avatars/1.jpg');

INSERT INTO recipes (id, title, cooktime, calories, rating, title_image_url, author_id, description) VALUES (1, 'Новогодний очень вкусный салат', 235, 100, 5, 'https://yaroslavs-imac.local:80/static/recipes/салат.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1', 'Например, очень вкусный салат, который готовили народы майя перед проведением жертвоприношений.'), (2, 'Тыквенный суп котик мяу мяу', 55, 110, 4, 'https://yaroslavs-imac.local:80/static/recipes/тыкв_суп.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1', 'Ну, допустим, мяу.');

INSERT INTO recipe_compilations VALUES (1, 'Аюрведические блюда');
INSERT INTO recipes (id, title, cooktime, calories, rating, title_image_url, author_id) VALUES (3, 'Кичари', 70, 50, 5, 'https://yaroslavs-imac.local:80/static/recipes/кичари.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1', 'Могучая аюрведическая штука, с имбиря неплохо штырит, рекомендую. На удивление хорошо сочетаются рис и маш, считаю, стоит того, чтобы попробовать.'), (4, 'Чечевичные котлеты', 60, 90, 3, 'https://yaroslavs-imac.local:80/static/recipes/чечевич_котлеты.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1', 'Котлеты для правильного питания с неправильным вкусом. Если котлеты, то в них должно быть мясо, но тут только здоровая и полезная чечевица, от которой слегка сводит челюсть. Но это же не остановит тебя в стремлении похудеть?');
INSERT INTO recipe_compilations_recipes VALUES (1, 3, 1), (2, 4, 1);

INSERT INTO recipes (id, title, cooktime, calories, rating, title_image_url, author_id) VALUES (5, 'Перловый суп с морским вкусом', 160, 80, 5, 'https://yaroslavs-imac.local:80/static/recipes/перл_суп.jpeg', 'WPLfq0KRbMg7ziylPtR3d2Wditr1', 'Человек укусил пчелу, и она разбухла. Человек полетел собирать мёд, а пчела пошла на работу. Круговорот человека и пчелы...');
INSERT INTO promo_recipes VALUES (1, 5);

INSERT INTO units VALUES (1, 'г'), (2, 'мл'), (3, 'шт');

INSERT INTO ingredients VALUES (1, 'Картофель'), (2, 'Морковь'), (3, 'Вода'), (4, 'Рис'), (5, 'Имбирь');

INSERT INTO recipes_ingredients VALUES (1, 2, 1, 1, 400), (2, 2, 2, 1, 300), (3, 2, 3, 2, 1500);
INSERT INTO recipes_ingredients VALUES (4, 1, 1, 3, 2);
INSERT INTO recipes_ingredients VALUES (5, 5, 1, 1, 400), (6, 5, 3, 2, 2000), (7, 5, 4, 1, 350), (8, 5, 5, 3, 2);

INSERT INTO recipe_medias VALUES (2, 'https://static.sobaka.ru/images/image/01/63/21/50/_normal.jpeg?v=1671365310', 1), (2, 'https://ir.ozone.ru/s3/multimedia-m/c1000/6275130526.jpg', 1);
INSERT INTO recipe_medias VALUES (1, 'http://jplayer.org/video/m4v/Finding_Nemo_Teaser.m4v', 2);
INSERT INTO recipe_medias VALUES (5, 'https://www.purina.ru/sites/default/files/styles/nppe_breed_selector_500/public/2020-04/french_bulldog.jpg?itok=CWfOzyGk', 1), (5, 'http://jplayer.org/video/m4v/Incredibles_Teaser_320x144_h264aac.m4v', 2), (5, 'https://www.bethowen.ru/upload/medialibrary/ecc/eccb9676c5dbd1524849166d27627e24.jpg', 1);

INSERT INTO recipes_steps VALUES (1, 1, 'Картошка', 'Нарежь картоху для парилки', 2), (2, 2, 'Варка', 'Закинь вариться, пока не станет мягкой', 2), (3, 3, 'Финалочка', 'Произвольно намешай еще чего-нибудь, чтобы было вкусно', 2);
INSERT INTO recipes_steps VALUES (4, 1, 'Краткость - сестра таланта, как в свое время писал Антон Павлович Чехов', 'Рецепт с одним пунктом, чтобы был', 4);
INSERT INTO recipes_steps VALUES (5, 1, 'Пункт 1', 'Мяу', 5), (6, 2, 'Пункт 2', 'Длинный-длинный текст, чтобы проверить многострочное отображение пункта приготовления пищи для поедания и удовлетворения примитивных потребностей организма, продлевающих бытие в этом чудесном мире', 5), (7, 3, 'Пункт 3', 'Мяу-мяу', 5);

INSERT INTO recipe_food_values VALUES (1, 10, 10, 10, 1), (2, 20, 20, 20, 2), (3, 50, 150, 10, 3), (4, 200, 20, 220, 4), (5, 50, 50, 5, 5); 