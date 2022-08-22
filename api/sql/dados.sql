insert into usuarios(nome, nick, email, senha)
values
('Usuario 1', 'usuario_1', 'user1@gmail.com', '$2a$10$qukgymNNTUHlLgC4TSqUz.MSNdI3EQzvrRnk3Fqzf4Yo.tsWwPqb6'),
('Usuario 2', 'usuario_2', 'user2@gmail.com', '$2a$10$qukgymNNTUHlLgC4TSqUz.MSNdI3EQzvrRnk3Fqzf4Yo.tsWwPqb6'),
('Usuario 3', 'usuario_3', 'user3@gmail.com', '$2a$10$qukgymNNTUHlLgC4TSqUz.MSNdI3EQzvrRnk3Fqzf4Yo.tsWwPqb6'),
('Usuario 4', 'usuario_4', 'user4@gmail.com', '$2a$10$qukgymNNTUHlLgC4TSqUz.MSNdI3EQzvrRnk3Fqzf4Yo.tsWwPqb6');


insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3) ON CONFLICT (usuario_id, seguidor_id) DO NOTHING;