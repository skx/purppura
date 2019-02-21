# First of all create the database:
#
#   CREATE DATABASE purple;
#
# Now create a user to access it:
#
#   CREATE USER 'purple'@'%' IDENTIFIED BY 'purple';
#   GRANT ALL PRIVILEGES ON purple.* TO 'purple'@'%';
#
# Finally we can create the tables like so:
#


#
# Create the table for storing usernames & (bcrypt) password hashes.
#
CREATE TABLE users (
  i INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username char(65),
  password char(65)
);


#
# Create the table for events.
#
CREATE TABLE events (
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
