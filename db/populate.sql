INSERT INTO role (`name`, `description`)
VALUES ('admin', 'administrator'),('client', 'standard client');

INSERT INTO user (`name`, `email`, `password`, `blocked`, `role_id`)

VALUES ('admin', 'admin@sup-burger.com', '123456', 0, 1), ('dodo', 'dodo@sup-burger.com', '123456', 0,  2), ('John Doe', 'john@sup-burger.com', '123456', 1, 2);
