# ************************************************************
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
# ************************************************************
# Every table has created_at & updated_at 
# `created_at` timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
# `updated_at` timestamp    NULL DEFAULT CURRENT_TIMESTAMP,

CREATE TABLE `goadmin_menu` (
  `id`         int(10)      NOT NULL AUTO_INCREMENT,
  `parent_id`  int(11)      NOT NULL DEFAULT '0',
  `type`       tinyint(4)   NOT NULL DEFAULT '0',
  `order`      int(11)      NOT NULL DEFAULT '0',
  `title`      varchar(50)  NOT NULL,
  `icon`       varchar(50)  NOT NULL,
  `uri`        varchar(50)  NOT NULL DEFAULT '',
  `header`     varchar(150) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `goadmin_menu` (`id`, `parent_id`, `type`, `order`, `title`, `icon`, `uri`, `header`)
VALUES
	(1, 0, 1, 2, 'Admin',         'fa-tasks',     '',                 NULL),
	(2, 1, 1, 2, 'Users',         'fa-users',     '/info/manager',    NULL),
	(3, 1, 1, 3, 'Roles',         'fa-user',      '/info/roles',      NULL),
	(4, 1, 1, 4, 'Permission',    'fa-ban',       '/info/permission', NULL),
	(5, 1, 1, 5, 'Menu',          'fa-bars',      '/menu',            NULL),
	(6, 1, 1, 6, 'Operation log', 'fa-history',   '/info/op',         NULL),
	(7, 0, 1, 1, 'Dashboard',     'fa-bar-chart', '/',                NULL);

CREATE TABLE `goadmin_operation_log` (
  `id`         int(10)      NOT NULL AUTO_INCREMENT,
  `user_id`    int(11)      NOT NULL,
  `path`       varchar(255) NOT NULL,
  `method`     varchar(10)  NOT NULL,
  `ip`         varchar(15)  NOT NULL,
  `input`      text         NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `goadmin_permissions` (
  `id`          int(10)      NOT NULL AUTO_INCREMENT,
  `name`        varchar(50)  NOT NULL,
  `slug`        varchar(50)  NOT NULL,
  `http_method` varchar(255) DEFAULT NULL,
  `http_path`   text         NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `goadmin_permissions` (`id`, `name`, `slug`, `http_method`, `http_path`) VALUES
	(1, 'All permission', '*',         '',                    '*'),
	(2, 'Dashboard',      'dashboard', 'GET,PUT,POST,DELETE', '/');

CREATE TABLE `goadmin_role_menu` (
  `role_id` int(11) NOT NULL,
  `menu_id` int(11) NOT NULL,
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `goadmin_role_menu` (`role_id`, `menu_id`) VALUES
	( 1,  1 ), ( 1,  7 ), ( 2,  7 ), ( 1,  8 ), ( 2,  8 );

CREATE TABLE `goadmin_role_permissions` (
  `role_id`       int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  UNIQUE KEY `admin_role_permissions` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `goadmin_role_permissions` (`role_id`, `permission_id`) VALUES
	( 1, 1 ), ( 1, 2 ), ( 2, 2 );

CREATE TABLE `goadmin_role_users` (
  `role_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  UNIQUE KEY `admin_user_roles` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `goadmin_role_users` (`role_id`, `user_id`) VALUES ( 1, 1 ), ( 2, 2 );

CREATE TABLE `goadmin_roles` (
  `id`         int(10)     NOT NULL AUTO_INCREMENT,
  `name`       varchar(50) NOT NULL,
  `slug`       varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `goadmin_roles` (`id`, `name`, `slug`) VALUES
	(1, 'Administrator', 'administrator'),
	(2, 'Operator',      'operator');

CREATE TABLE `goadmin_session` (
  `id`         int(11)       NOT NULL AUTO_INCREMENT,
  `sid`        varchar(50)   NOT NULL DEFAULT '',
  `values`     varchar(3000) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `goadmin_user_permissions` (
  `user_id`       int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  UNIQUE KEY `admin_user_permissions` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `goadmin_user_permissions` (`user_id`, `permission_id`) VALUES
	( 1, 1 ), ( 2, 2 );

CREATE TABLE `goadmin_users` (
  `id`             int(10)      NOT NULL AUTO_INCREMENT,
  `username`       varchar(190) NOT NULL,
  `password`       varchar(80)  NOT NULL,
  `name`           varchar(255) NOT NULL,
  `avatar`         varchar(255) DEFAULT NULL,
  `remember_token` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_unique` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
