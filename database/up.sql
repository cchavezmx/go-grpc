DROP TABLE IF EXISTS `users`;
CREATE TABLE students (
  id VARCHAR(32) PRIMARY KEY,
  NAME VARCHAR(255) NOT NULL,
  age INTEGER NOT NULL,
);