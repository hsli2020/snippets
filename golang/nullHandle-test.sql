use test;
CREATE TABLE Article(
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(100) NOT NULL,
    `pub_date` datetime DEFAULT NULL,
    `body` text,
    `user_id` int DEFAULT NULL,
    PRIMARY KEY(id)
);
INSERT INTO Article(`title`) VALUES("first article");