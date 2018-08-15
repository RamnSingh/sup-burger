INSERT INTO role (`name`, `description`)
VALUES ('admin', 'administrator'),('client', 'standard client');
INSERT INTO user (`name`, `email`, `blocked`, `img_path`, `role_id`)
VALUES ('admin', 'admin@sup-burger.com', 0, 'users\\admin.jpg', 1), ('dodo', 'dodo@sup-burger.com', 0, 'users\\user.jpg', 2), ('John Doe', 'john@sup-burger.com', 1, 'users\\user_1.jpg', 2);
