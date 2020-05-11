# This is a simple SQL file which is designed to create a new user,
# and a new database.
#
# If you're running this manually you can create the user, password, and
# database like so:
#
#
#   CREATE DATABASE purple;
#
# Now create a user to access it:
#
#   CREATE USER 'purple'@'%' IDENTIFIED BY 'purple';
#   GRANT ALL PRIVILEGES ON purple.* TO 'purple'@'%';
#
#
# This file is executed by Docker, so it creates a fixed user but you can
# skip that if you jump over the following entries.
#


#
# Create the database
#
CREATE DATABASE IF NOT EXISTS purple;

#
# Ensure we have a purple-username/password
#
CREATE USER IF NOT EXISTS 'purple'@'%' IDENTIFIED BY 'purple';
GRANT ALL ON `purple`.* TO 'purple'@'%';
FLUSH PRIVILEGES;

USE purple;


#
# Create the table for storing usernames & (bcrypt) password hashes.
#
CREATE TABLE IF NOT EXISTS users (
  i INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username char(65),
  password char(65)
);


#
# Create the table for events.
#
CREATE TABLE IF NOT EXISTS events (
  i BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  id     char(100) not null,
  source char(100) not null,
  status char(15) DEFAULT 'pending',
  raise_at int default '0',
  notified_at int default '0',
  notify_count int default '0',
  subject text not null,
  detail  text not null
);

# Add an index here, for querying.
ALTER TABLE events ADD INDEX find (id, source);
