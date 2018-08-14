INSERT INTO role (`name`, `description`)
VALUES ('admin', 'administrator'),('client', 'standard client');
INSERT INTO user (`name`, `blocked`, `img_path`, `role_id`)
VALUES ('admin', 0, 'users\\admin.jpg', 1), ('dodo', 0, 'users\\user.jpg', 2), ('John Doe', 1, 'users\\user_1.jpg', 2);
