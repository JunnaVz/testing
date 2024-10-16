--1. create tables/extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists users cascade;
create table users
(
    id           uuid primary key default uuid_generate_v4(),
    name         text,
    surname      text,
    email        text unique,
    phone_number text,
    address      text,
    password     text
);
ALTER TABLE users
    ALTER COLUMN id SET DEFAULT uuid_generate_v4();

drop table if exists workers cascade;
create table workers
(
    id           uuid primary key default uuid_generate_v4(),
    name         text,
    surname      text,
    email        text unique,
    phone_number text,
    address      text,
    password     text,
    role         int
);

ALTER TABLE workers
    ALTER COLUMN id SET DEFAULT uuid_generate_v4();

drop table if exists orders cascade;
create table orders
(
    id            uuid primary key                                default uuid_generate_v4(),
    worker_id     uuid references workers (id) on delete set null default null,
    user_id       uuid references users (id) on delete set null   default null,
    status        int2                                            default 0,
    address       text,
    deadline      timestamp,
    creation_date timestamp                                       default now(),
    rate          int2                                            default 0
);
ALTER TABLE orders
    ALTER COLUMN id SET DEFAULT uuid_generate_v4(),
    ALTER COLUMN creation_date SET DEFAULT now(),
    ALTER COLUMN status SET DEFAULT 0,
    ALTER COLUMN rate SET DEFAULT 0;


drop table if exists tasks cascade;
create table tasks
(
    id               uuid primary key default uuid_generate_v4(),
    name             text,
    price_per_single float8,
    category         int2
);
ALTER TABLE tasks
    ALTER COLUMN id SET DEFAULT uuid_generate_v4();

drop table if exists order_contains_tasks cascade;
create table order_contains_tasks
(
    id       uuid primary key default uuid_generate_v4(),
    order_id uuid references orders (id),
    task_id  uuid references tasks (id),
    quantity int2             default 1
);
ALTER TABLE order_contains_tasks
    ALTER COLUMN id SET DEFAULT uuid_generate_v4();

drop table if exists categories cascade;
CREATE TABLE IF NOT EXISTS categories
(
    id   SERIAL UNIQUE,
    name VARCHAR
);
SELECT setval(pg_get_serial_sequence('categories', 'id'), COALESCE(MAX(id), 1) + 1, false)
FROM categories;

INSERT INTO categories (id, name)
VALUES (1, 'Мытье окон'),
       (2, 'Мытье окон'),
       (3, 'Мытье окон'),
       (4, 'Мытье окон'),
       (5, 'Мытье окон'),
       (6, 'Мытье окон'),
       (7, 'Мытье окон'),
       (8, 'Мытье окон');

--2. insert into
INSERT INTO tasks (id, name, price_per_single, category)
VALUES ('f7f4962c-d2a6-4d30-bea4-c928f7642e94', 'Шторы из плотной ткани', '200', '8'),
       ('949b73f5-7270-4c8f-82e6-6f2907b1e1d6', 'Ламбрекен', '200', '8'),
       ('09807123-39d2-468c-922d-675262add44b', 'Драпировка настенная', '200', '8'),
       ('ae2a99c0-c4f7-4b6c-a4ca-9f857a593bc7', 'Штора рулонаая', '200', '8'),
       ('9b205a67-2b03-4418-97cd-7dd5ccf8208a', 'Жалюзи вертикальные', '200', '8'),
       ('d85c47d3-63de-4db7-b79d-a7b1aaf3c4aa', 'Стул', '200', '8'),
       ('e8d6561b-cb25-4a86-a5e5-d2ce0e62b6a6', 'Пуфик', '250', '8'),
       ('26034dfd-a91a-47db-848a-1d379869a6d3', 'Компьютерное кресло', '300', '8'),
       ('b83e962c-2d7c-47e6-97d8-fb66aad0a088', 'Кресло', '800', '8'),
       ('3f25767e-b0d9-4795-9184-87556bb74b09', 'Диванная подушка', '200', '8'),
       ('8ce932b7-f871-4865-9738-690674822130', 'Диван (2-х местный)', '1200', '8'),
       ('00bedf04-5f05-4e4f-a7ef-01e793b3d8f7', 'Диван (3-х местный)', '1500', '8'),
       ('5d0624ec-6d42-4f80-8766-826661a9658a', 'Диван (4-х местный)', '1800', '8'),
       ('1ff5ac54-e4e9-4de1-afb2-b8fc397dcead', 'Диван (5-х местный)', '2100', '8'),
       ('af6c07e1-9b37-4f2d-b492-2f4ebeeb26ee', 'Матрас детский', '600', '8'),
       ('5e672378-ef9a-478b-b238-29e70a99d455', 'Матрас односпальный', '1000', '8'),
       ('3269eeab-07e0-4a2b-bbd4-51cbb57c7e8b', 'Матрас полутороспальный', '1400', '8'),
       ('e060b4e3-e5d1-457b-8d17-c8069ac77e9b', 'Двуспальный матрас', '2000', '8'),
       ('1079c7bd-3a0c-4df2-810c-b64d293eebae', 'Одеяло / покрывало', '200', '8'),
       ('7f7a4dfb-c158-4a1a-a2d4-45b3fcafc6b4', 'Подушки спальные', '200', '8'),
       ('680d28f1-9161-4359-b7e9-7179a8cfc768', 'Ковролин / ковровая плитка', '140', '8');


INSERT INTO tasks (id, name, price_per_single, category)
VALUES ('8d9e0a5a-7a3d-42c7-a916-35bcfba5b9b2', 'Сухая и влажная уборка твёрдых полов', '200', '7'),
       ('6d407e05-2876-44ff-9d10-83e32a2fba52', 'Глубокая очистка полов роторной машиной', '300', '7'),
       ('5487c850-f658-4580-a471-d6c7dd3cf0a7', 'Шлифовка', '200', '7'),
       ('cd558d78-c34e-47dd-8079-602839cca2e0', 'Полировка', '200', '7'),
       ('f14b01d9-abe5-4fb2-98da-dc38effadf53', 'Кристаллизация', '400', '7'),
       ('ad012253-2ee7-42c8-a1ad-471334746938', 'Нанесение защитного покрытия', '200', '7');

INSERT INTO tasks (id, name, price_per_single, category)
VALUES ('bf96d4e3-3ca0-4d1b-8058-570be8bdd87f', 'Стул', '500', '6'),
       ('346c07d3-e81c-43e5-961e-5876648db2e0', 'Пуфик', '600', '6'),
       ('9c94ec24-4ab5-49e1-a423-fc9b5793f3af', 'Компьтерное кресло', '750', '6'),
       ('bb2dc461-4d1b-48b9-bea6-b65c7e6e0e52', 'Кресло', '1900', '6'),
       ('d081fa3f-8d14-45d6-a4d7-93f5bdf64a77', 'Диванная подушка', '500', '6'),
       ('26031e20-48d5-4085-bf9c-7a0ee926aa9d', 'Диван (2-х местный)', '2800', '6'),
       ('de9fdfe4-350e-4d51-b8b7-34097862a4d7', 'Диван (3-х местный)', '3600', '6'),
       ('1955c0d3-e394-4f09-a1ee-3e7afd5afe00', 'Диван (4-х местный)', '4600', '6'),
       ('117cb945-b94b-4f4e-a719-642f1684cda9', 'Диван (5-х местный)', '5600', '6'),
       ('958fd106-2beb-423c-a32f-444a9d4a9e27', 'Ковролин / ковровая плитка', '200', '6');


INSERT INTO tasks (id, name, price_per_single, category)
VALUES ('51ae888b-264d-4e80-8aea-bd27c18ebc6b', 'Удаление пыли и загрязнений', '500', '5'),
       ('53dc1708-bd2a-4cb1-b5d6-62049167432f', 'Сухая уборка пола, чистка ковролина и напольных ковров пылесосом',
        '500', '5'),
       ('4e39eca6-462e-4fd1-9499-2e05d1d02133',
        'Влажная уборка пола и плинтусов с помощью специальных ухаживающих средств', '500', '5'),
       ('7e8028cd-554c-4bf5-a8f3-5fb6e3f12cf1', 'Сухая уборка пылесосом мягкой мебели', '1000', '5'),
       ('b3b2a956-3169-48ad-8db7-33536dd80cca', 'Влажная уборка подоконников, отопительных труб, радиаторов', '500',
        '5'),
       ('52444b8c-3aac-4e2a-9790-a68db0028206', 'Влажная уборка дверей, наличников, дверной фурнитуры', '200', '5'),
       ('d9d1dda8-dfa9-476d-86b5-77de770aff14', 'Мытьё рабочей поверхности кухонной плиты снаружи', '200', '5'),
       ('e8247504-9924-4012-b915-45be412f07db', 'Мытьё духовой печи снаружи', '500', '5'),
       ('db25e331-c64f-448d-831c-2e4a365659a9', 'Мытьё рабочих поверхностей (столешниц, барных стоек', '500', '5'),
       ('58c3b2ed-22b7-4db2-b40d-38cdfdf6fb0a', 'Мытьё кухонного фартука', '500', '5'),
       ('9dd23974-ab63-4eff-ab75-a38f1eed098f', 'Мытьё внешних вертикальных поверхностей кухонных шкафов', '500', '5'),
       ('0192b034-2d68-4473-a1d5-bd50d23e25bc', 'Мытьё раковины', '100', '5'),
       ('0585d17e-708c-4cd8-84c4-93467a289d7e',
        'Влажная уборка полов / плинтусов с помощью специальных ухаживающих средств ', '500', '5'),
       ('dc9a0d5e-376e-4475-abcb-f157d9fce268', 'Удаление пыли / загрязнений вытяжки ', '200', '5'),
       ('984895e3-da71-44b3-b439-f93e7038c8a8', 'Влажная уборка мест хранения мусора ', '200', '5'),
       ('3f630f02-0bc9-4253-82c6-0f0f7be581c1', 'Мытьё сантехники', '500', '5'),
       ('788d5f47-a3e4-44f1-b0e9-a8f5a45e1f1b', 'Мытьё стен', '500', '5'),
       ('0d9cd6b1-1323-461b-92cb-784c9d0ed680',
        'Очистка внешней поверхности зеркальных шкафов, стиральной и сушильной машины ', '500', '5');


INSERT INTO tasks (id, name, price_per_single, category)
VALUES ('75a99a0c-5bb1-447a-9201-baf6f5c9b4a8',
        'Влажная очистка потолочных и настенных светильников от пыли и загрязнений', '500', '4'),
       ('b8813256-63a6-4013-9913-891686c6a907', 'Сухая и влажная уборка влагостойкого поктрытия стен', '500', '4'),
       ('a9e0d89e-340e-400b-a789-cb7687ede514', 'Сухая уборка потолков', '200', '4'),
       ('63a10915-e98b-4598-a212-06858ff206b4',
        'Удаление загрязнений с входных и межкомнатных дверей, дверных коробок, доводчиков', '500', '4'),
       ('1e3af184-1a54-46de-8fb7-e5ed2b350d03', 'Влажная уборка вертикальных и горизонтальных поверхностей мебели',
        '200', '4'),
       ('d925b987-a4be-4a61-aa5c-75914e144287', 'Сухая и влажная уборка пола / плинтусов', '500', '4'),
       ('001264f9-e389-4964-98d1-a0379222c406',
        'Влажная очистка от пыли пожарных шкафов, батарей и труб центрального отопления, решеток вентиляции', '200',
        '4'),
       ('c3c71344-fe9d-46b1-a1e9-08f1aa8fa62c',
        'Влажная очистка от пыли и загрязнений элементов декора, картин, пластиковых коробов', '150', '4'),
       ('8dfe0f7c-4928-4dac-a07d-41e1ca767087', 'Влажная уборка рабочих столов, протирка оргтехники, крестовин кресел',
        '150', '4'),
       ('a58a5d98-f862-4919-a022-82f83d44871a',
        'Сбор и вынос мусора из мусорных корзин, протирка мусорных корзин, замена полиэтиленовых пакетов', '100', '4'),
       ('0058e0e0-c143-42a7-9098-204fb707cf89', 'Влажная очистка подоконников от пыли и загрязнений', '200', '4'),
       ('437a91cc-95ea-4b29-9d02-95426c5ded6e',
        'Влажная очистка от пыли и загрязнений зеркальных и стеклянных поверхностей', '200', '4'),
       ('afc96510-b309-4f43-b9da-5367466c4e48',
        'Влажная уборка помещений для приема пищи, чистка кухонной техники, мытье кулеров для воды', '200', '4'),
       ('e29c4547-b522-45f9-9c4b-75c0bb11924d',
        'Комплексная и поддерживающая уборка санузлов, замена расходных материалов', '200', '4');

INSERT INTO tasks (id, name, price_per_single, category)
VALUES ('421b342c-ee23-4881-bdd4-739ec7688e35', 'Мытьё стекол', '350', '3'),
       ('c1a7f1a8-5225-4250-8ee8-e989195b45ac', 'Мытьё оконных рам', '250', '3'),
       ('f641c41a-cb3f-4096-961a-362097da3200', 'Мытьё откосов', '200', '3'),
       ('56aae491-b935-4f7c-98fa-2a28842ad394', 'Мытьё подоконников', '200', '3'),
       ('facccee6-fdfb-4cb2-bad2-68e5b3576500', 'Мытьё слива снаружи', '200', '3'),
       ('bd99fecb-8ae5-4461-a064-df7e207b264c', 'Мытьё москитных сеток', '200', '3'),
       ('873e2351-75ba-4c09-a6dc-d472d33bf5d6', 'Мытьё оконных решеток', '500', '3'),
       ('5917cd95-564b-4537-8d16-9e9bf34ef7f0', 'Снять / повесить шторы', '500', '3');


INSERT INTO tasks (id, name, price_per_single, category)
VALUES ('e523ff3e-d814-4bce-b0a2-eb7e4ebc35b7', 'Сбор и вынос мелкого строительного мусора', '500', '2'),
       ('f16a133d-1649-42c7-b89d-2b17f34b8545', 'Удаление строительной пыли сухим способом с потолка, стен, мебели',
        '500', '2'),
       ('141a4f19-24b7-4066-b197-fe8d0ac4b8b0',
        'Влажная уборка всех поверхностей, осветительных и отопительных приборов, дверей, подоконников и плинтусов',
        '500', '2'),
       ('67dc17ed-a94d-4acf-a29c-6771584376b0', 'Удаление локальных строительных загрязнений со всех поверхностей',
        '200', '2'),
       ('0700ee1d-456a-4641-909e-7b72d7c7d1c0', 'Влажная уборка внутренних поверхностей корпусной мебели', '200', '2'),
       ('1a6f84bf-1eb3-4d0c-9444-44cdef1c1c6c', 'Сухая чистка пылесосом мягкой мебели снаружи и внутри', '200', '2'),
       ('eb0afcdf-0f3a-4b89-bc0a-6a8acc7b9e83', 'Комплексная уборка санузлов, мытьё и дезинфекция сантехники', '300',
        '2'),
       ('91b5ea8e-36ce-47d4-a35a-d6f0e61e5a8a', 'Мытьё стеклянных и зеркальных поверхностей внутри помещения', '300',
        '2');

INSERT INTO tasks (id, name, price_per_single, category)
VALUES ('daa09f13-0ba4-4511-a105-0e612ca11603', 'Обеспыливание стен, потолков, карнизов и кондиционеров', '300', '1'),
       ('3068fe74-e9fc-40ac-9674-e0bef4f83083',
        'Удаление известкового налетка, ржавйины, жира, водного камня со всех твердых поверхностей', '500', '1'),
       ('495fba4a-1d12-48bf-8292-5e357862c718', 'Чистка парогенератором труднодоступных мест', '400', '1'),
       ('3e9d5489-2ff0-4145-80e6-78f8b6c07161', 'Протирка настенных объектов - розетки и выключатели', '100', '1'),
       ('b46d01fb-307d-43a9-a048-af5f9d0ceea3',
        'Удаление пыли и загрязнений с веншних поверхностей мебели, бытовой техники, крупных предметов интерьера и домашнего декора',
        '300', '1'),
       ('591e9845-c064-496b-bb98-a78e3d6c724d', 'Влажная уборка внутренних поверхностей корпусной мебели', '300', '1'),
       ('42575ec6-de5d-422d-92bd-a46f1c9219c7', 'Очистка кафельной плитки, очистка межплиточных швов', '200', '1'),
       ('c14fdacf-e08e-4840-b9c8-25ebcc17df7e', 'Мытьё пластиковых / рееечных потолков', '300', '1'),
       ('7a5fdbf8-c658-4085-9754-4280c919cdef', 'Удаление загрязнений с осветительных приборов', '300', '1'),
       ('ae213344-27c1-4db6-839a-9ff4e59d389b', 'Протирка зеркал, стеклянных и отражающих поверхностей', '300', '1'),
       ('a51ee158-4c76-404a-9b08-d73e9b1a1b40', 'Влажная уборка подоконников, отопительных труб, радиаторов, экранов',
        '200', '1'),
       ('1a940e45-0b66-4604-bab1-2bc4ddd04623',
        'Чистка пылесосом мягкой мебели снаружи и внутри, ковров и ковровых покрытий', '200', '1'),
       ('1eff4b81-1254-492d-b33e-99e32b19aa96',
        'Сухая и влажная уборка полов и плинтусов с помощью бактерицидных и специальных ухаживающих средств', '350',
        '1'),
       ('7b132974-038a-451b-a254-61461c177266', 'Мытьё и удаление жира с бытовой техники снаружи', '500', '1'),
       ('a092cc99-5362-47f3-85d8-9279fe411ec7', 'Очистка и дезинфекция сантехники ', '500', '1'),
       ('240c7cd1-c130-46e0-b968-1726b4c89d63', 'Протирка дверей, наличников, дверной фурнитуры ', '200', '1'),
       ('f0b892b6-56c5-4051-8b9d-966c5fab1414', 'Влажная уборка и дезинфекция мест хранения мусора ', '300', '1');

--3. read files with data
COPY public.workers (ID, NAME, SURNAME, EMAIL, PHONE_NUMBER, ADDRESS, PASSWORD, ROLE)
    FROM '/tmp/workers_data.csv' DELIMITER ';' CSV HEADER NULL 'NULL';

COPY public.users (ID, NAME, SURNAME, EMAIL, PHONE_NUMBER, ADDRESS, PASSWORD)
    FROM '/tmp/users_data.csv' DELIMITER ';' CSV HEADER NULL 'NULL';

COPY public.orders (ID, WORKER_ID, USER_ID, STATUS, DEADLINE, ADDRESS, CREATION_DATE, RATE)
    FROM '/tmp/orders_data.csv' DELIMITER ';' CSV HEADER NULL 'NULL';

COPY public.order_contains_tasks (ID, ORDER_ID, TASK_ID, QUANTITY)
    FROM '/tmp/order_contains_data.csv' DELIMITER ';' CSV HEADER NULL 'NULL';

--4. delete
DELETE
FROM workers
WHERE role = 1;

UPDATE workers
SET password = '$2b$12$TbfG11CRR9OSEsNX.Awije1.DmStMp.Erq1nJ/xIYwu.ilYjSbwOm';

UPDATE users
SET password = '$2b$12$TbfG11CRR9OSEsNX.Awije1.DmStMp.Erq1nJ/xIYwu.ilYjSbwOm';