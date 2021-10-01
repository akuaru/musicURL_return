CREATE DATABASE IF NOT EXISTS otokake;
USE otokake;
CREATE TABLE musics(
music_ID INT NOT NULL AUTO_INCREMENT UNIQUE,
music_URL VARCHAR(45) NOT NULL UNIQUE,
music_artist VARCHAR(45) NOT NULL,
music_name VARCHAR(45) NOT NULL,
PRIMARY KEY (music_ID));
SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;

INSERT INTO musics (music_URL,music_artist,music_name) VALUES
("http://hoge/otokake/sample01","hoge","sample01"),
("http://hoge/otokake/sample02","hoge","さんぷる02"),
("http://hoge/otokake/sample03","hoge","サンプル03"),
("http://hoge/otokake/sample04","hoge","ｻﾝﾌﾟﾙ04"),
("http://hoge/otokake/sample05","hoge","sample05"),
("http://hoge/otokake/sample06","fuga","さんぷる06"),
("http://hoge/otokake/sample07","fuga","サンプル07"),
("http://hoge/otokake/sample08","fuga","ｻﾝﾌﾟﾙ08"),
("http://hoge/otokake/sample09","foo","sample09"),
("http://hoge/otokake/sample10","foo","さんぷる10");