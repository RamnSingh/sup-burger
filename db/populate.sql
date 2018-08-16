INSERT INTO role (`name`, `description`)
VALUES ('admin', 'administrator'),('client', 'standard client');
INSERT INTO user (`name`, `email`, `password`, `blocked`, `img_path`, `role_id`)
VALUES ('admin', 'admin@sup-burger.com', '123456', 0, 'users\\admin.jpg', 1), ('dodo', 'dodo@sup-burger.com', '123456', 0, 'users\\user.jpg', 2), ('John Doe', 'john@sup-burger.com', '123456', 1, 'users\\user_1.jpg', 2);
