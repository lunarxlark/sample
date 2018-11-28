CREATE USER tech_app IDENTIFIED BY 'password';
GRANT ALL ON teck_link.* TO 'tech_app'@'%';

CREATE TABLE language (
  name VARCHAR(20),
  version VARCHAR(20),
  release_date DATE
);
