LOAD DATA INFILE '/var/lib/mysql-files/golang.csv'
  INTO TABLE tech_link.language
  FIELDS
    TERMINATED BY ','
    ENCLOSED BY '"'
