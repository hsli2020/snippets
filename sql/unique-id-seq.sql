-- https://blog.huoding.com/2019/08/21/768

-- 方案一
mysql> CREATE TABLE `Tickets64` (
  `id` bigint(20) unsigned NOT NULL auto_increment,
  `stub` char(1) NOT NULL default '',
  PRIMARY KEY  (`id`),
  UNIQUE KEY `stub` (`stub`)
) ENGINE=MyISAM;

mysql> REPLACE INTO Tickets64 (stub) VALUES ('a');
mysql> SELECT LAST_INSERT_ID();

-- 方案二
mysql> CREATE TABLE `seq` (
  `id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  UNIQUE KEY `name` (`name`)
) ENGINE=MyISAM;

mysql> INSERT INTO seq (id, name) VALUES (0, 'global');
mysql> INSERT INTO seq (id, name) VALUES (0, 'another');

mysql> UPDATE seq SET id = LAST_INSERT_ID(id+1) WHERE name = 'global';
mysql> SELECT LAST_INSERT_ID();