
-- u-sushi hash
-- $argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4

truncate table sushi_user restart identity cascade;
truncate table plate_ingredient restart identity cascade;
truncate table ingredient restart identity cascade;
truncate table allergen restart identity cascade;
truncate table plate restart identity cascade;
truncate table category restart identity cascade;
truncate table image restart identity cascade;

insert into image (image) values 
('test-image-1'),
('test-image-2');

insert into category (name) values 
('test-category-1');

insert into plate (name, price, category_id, menu, pieces) values
('test-plate-1', '4.0', 1, 'Lunch', 2),
('test-plate-2', '4.0', 1, 'Lunch', 2);

insert into allergen (name) values
('test-allergen-1');

insert into ingredient (name, allergen_id) values
('test-ingredient-1', 1),
('test-ingredient-2', 1);

insert into sushi_user (user_type, password) values 
('Admin', '$argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4'),
('Client', '$argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4'),
('Employee', '$argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4');