CREATE DATABASE purple;

CREATE USER 'purple'@'%' IDENTIFIED BY 'purple';
GRANT ALL PRIVILEGES ON purple.* TO 'purple'@'%';


CREATE TABLE users (
  i INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username char(65),
  password char(65)
);


CREATE TABLE events (
  i INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
  id    text not null,
  source text not null,
  status char(10) DEFAULT 'pending',
  raise_at int default '0',
  notified_at int default '0',
  subject text not null,
  detail  text not null        );
