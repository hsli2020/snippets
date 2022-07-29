<?php

# https://www.youtube.com/watch?v=K0bYK5zVglY

/*
SELECT
	users.id,
	users.name,
	GROUP_CONCAT(roles.name SEPARATOR ', ') AS role_names
FROM
	users
INNER JOIN role_user ON 
	role_user.user_id = users.id
INNER JOIN roles ON
	roles.id = role_user.role_id
GROUP BY
	users.id
*/


User::query()
	->select([
		'users.id',
		'users.name',
		DB::raw("GROUP_CONCAT(roles.name SEPARATOR ', ') AS role_names")
	])
	->join('role_user', 'role_user.user_id', '=', 'users.id')
	->join('roles', 'roles.id', '=', 'role_user.role_id')
	->groupBy('users.id')
	->get();
