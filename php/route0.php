<?php

# https://medium.com/@dylanbr/building-a-basic-router-b43c17361f8b

// Define routes
$routes = [
  '\/'             => function() { return '<a href="/hello">Click here for a greeting</a>'; },
  '\/hello'        => function() { return 'Hello world'; },
  '\/count\/(\d+)' => function($count) { return join("<br>", range(1,$count)); },
  '404'            => function() { return 'Page not found'; }
];

// Initialise the request, response and parameter list
if (PHP_SAPI == "cli") {
    $request = $argv[1];
} else {
    $request = $_SERVER['REQUEST_URI'];
}
$response = false;
$parameters = array();

// Find a route and define the response
foreach ($routes as $route=>$action) {
	if (preg_match('/^' . $route . '$/', $request, $matches)) {
		// throw away the full match at the start of the array
		$parameters = array_slice($matches, 1);
		$response = $action;
	}
}

// If no route was found, use the 404 route as the response
if ($response === false) {
	$response = $routes['404'];
}

// Output the response
echo call_user_func_array($response, $parameters), PHP_EOL;

/*
# Making string concatenation readable in PHP

1) bad
$logMessage = 'A '.$user->type.' with e-mail address '.$user->email.' has performed '.$action.' on '.$subject.'.';

2) good
$logMessage = sprintf('A %s with email %s has performed %s on %s.', $user->type, $user->email, $action, $subject);

3) better
$logMessage = "A {$user->type} with e-mail address {$user->email} has performed {$action} on {$subject}.";
*/
