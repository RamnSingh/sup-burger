DROP TABLE IF EXISTS 'stuff'
CREATE TABLE stuff (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) not null,
  description text not null,
  img_path varchar(1000) not null,
  primary key (id)
)ENGINE=InnoDB;


DROP TABLE IF EXISTS 'role'
CREATE TABLE role (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) not null,
  description text not null
)ENGINE=InnoDB;

DROP TABLE IF EXISTS 'city'
CREATE TABLE city (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) not null,
)ENGINE=InnoDB;


DROP TABLE IF EXISTS 'burger'
CREATE TABLE burger (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) not null,
  description text not null,
  img_path varchar(1000) not null,
  price decimal not null,
  stock smallint not null,
  stuff_id int not null,
  foreign key (stuff_id) references stuff(id)

)ENGINE=InnoDB;

DROP TABLE IF EXISTS 'burger_stuff'
CREATE TABLE burger_stuff (
  burger_id INT NOT NULL,
  stuff_id INT NOT NULL,
  foreign key (burger_id) references burger(id),
  foreign key (stuff_id) references stuff(id)

)ENGINE=InnoDB;


DROP TABLE IF EXISTS 'city_burger'
CREATE TABLE city_burger (
  burger_id INT NOT NULL,
  city_id INT NOT NULL,
  foreign key (burger_id) references burger(id),
  foreign key (city_id) references city(id)

)ENGINE=InnoDB;



DROP TABLE IF EXISTS 'order'
CREATE TABLE order (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  date datetime not null,
  bill_path varchar(1000) not null,
  total_price decimal not null,
  quantity smallint not null,
  burger_id int not null,
  user_id int not null,
  foreign key (stuff_id) references stuff(id)

)ENGINE=InnoDB;
