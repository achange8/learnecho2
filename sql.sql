create table boards
(
    NUM INT AUTO_INCREMENT PRIMARY KEY,
    TITLE VARCHAR(100) NOT NULL,
    WRITER VARCHAR(50) NOT NULL,
    CONTENT VARCHAR(4000) NOT NULL,
    DB_DATE varchar(10),
    HiTCOUNT INT default(0));

CREATE TABLE users(
    id varchar(30) NOT NULL UNIQUE,
    email VARCHAR(320) NOT NULL UNIQUE,
    password CHAR(60) NOT NULL);


    {
  "TITLE" : "tester04",
  "CONTENT":"test1234"
}

{
    "id" : "tester04",
    "password" : "test1234"
}