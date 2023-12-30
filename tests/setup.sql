
-- u-sushi hash
-- $argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4

truncate table plate cascade;
truncate table category cascade;

truncate table sushi_user;
insert into sushi_user (user_type, password) values 
('Admin', '$argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4'),
('Client', '$argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4'),
('Employee', '$argon2id$v=19$m=65536,t=3,p=4$vb4AULkt8aSGh/Rfq5BwOQ$jzmrQ/jv6I6e2MZhfejoIvUzaarV4874fgZ6cJS6eF4');