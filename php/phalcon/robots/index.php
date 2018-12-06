<?php

use Phalcon\Loader;
use Phalcon\Mvc\Micro;
use Phalcon\Di\FactoryDefault;
use Phalcon\Db\Adapter\Pdo\Mysql as PdoMysql;
use Phalcon\Http\Response;

// Use Loader() to autoload our model
$loader = new Loader();
$loader->registerNamespaces([ 'Store\Toys' => __DIR__ . '/models/' ]);
$loader->register();

$di = new FactoryDefault();

// Set up the database service
$di->set('db', function () {
    return new PdoMysql([
        'host'     => 'localhost',
        'username' => 'root',
        'password' => '',
        'dbname'   => 'phalconTutorial',
    ]);
});

// Create and bind the DI to the application
$app = new Micro($di);

// Define the routes here

$app->get('/api/robots', function () use ($app) {
		// Operation to fetch all the robots
    $phql = 'SELECT * FROM Store\Toys\Robots ORDER BY name';
    $robots = $app->modelsManager->executeQuery($phql);

    $data = [];
    foreach ($robots as $robot) {
        $data[] = [ 'id' => $robot->id, 'name' => $robot->name ];
    }

    echo json_encode($data);
});

// Searches for robots with $name in their name
$app->get('/api/robots/search/{name}', function ($name) use ($app) {
    // Operation to fetch robot with name $name
    $phql = 'SELECT * FROM Store\Toys\Robots WHERE name LIKE :name: ORDER BY name';
    $robots = $app->modelsManager->executeQuery($phql, [ 'name' => '%' . $name . '%' ]);

    $data = [];
    foreach ($robots as $robot) {
        $data[] = [ 'id' => $robot->id, 'name' => $robot->name ];
    }

    echo json_encode($data);
});

// Retrieves robots based on primary key
$app->get('/api/robots/{id:[0-9]+}', function ($id) use ($app) {
    // Operation to fetch robot with id $id
    $phql = 'SELECT * FROM Store\Toys\Robots WHERE id = :id:';
    $robot = $app->modelsManager->executeQuery($phql, [ 'id' => $id ])->getFirst();

    // Create a response
    $response = new Response();

    if ($robot === false) {
        $response->setJsonContent([ 'status' => 'NOT-FOUND' ]);
    } else {
        $response->setJsonContent([
            'status' => "FOUND",
            'data' => [ 'id' => $robot->id, 'name' => $robot->name ]
        ]);
    }

    return $response;
});

// Adds a new robot
$app->post('/api/robots', function () use ($app) {
    // Operation to create a fresh robot
    $robot = $app->request->getJsonRawBody();

    $phql = 'INSERT INTO Store\Toys\Robots (name, type, year) VALUES (:name:, :type:, :year:)';
    $status = $app->modelsManager->executeQuery($phql, [
        'name' => $robot->name,
        'type' => $robot->type,
        'year' => $robot->year,
    ]);

    // Create a response
    $response = new Response();

    // Check if the insertion was successful
    if ($status->success() === true) {
        $robot->id = $status->getModel()->id;

        // Change the HTTP status
        $response->setStatusCode(201, 'Created');
        $response->setJsonContent([ 'status' => 'OK', 'data' => $robot ]);
    } else {
        // Send errors to the client
        $errors = [];
        foreach ($status->getMessages() as $message) {
            $errors[] = $message->getMessage();
        }

        // Change the HTTP status
        $response->setStatusCode(409, 'Conflict');
        $response->setJsonContent([ 'status' => 'ERROR', 'messages' => $errors ]);
    }

    return $response;
});

// Updates robots based on primary key
$app->put('/api/robots/{id:[0-9]+}', function ($id) use ($app) {
    // Operation to update a robot with id $id
    $robot = $app->request->getJsonRawBody();

    $phql = 'UPDATE Store\Toys\Robots SET name = :name:, type = :type:, year = :year: WHERE id = :id:';
    $status = $app->modelsManager->executeQuery($phql, [
        'id' => $id,
        'name' => $robot->name,
        'type' => $robot->type,
        'year' => $robot->year,
    ]);

    // Create a response
    $response = new Response();

    // Check if the insertion was successful
    if ($status->success() === true) {
        $response->setJsonContent([ 'status' => 'OK' ]);
    } else {
        $errors = [];
        foreach ($status->getMessages() as $message) {
            $errors[] = $message->getMessage();
        }

        // Change the HTTP status
        $response->setStatusCode(409, 'Conflict');
        $response->setJsonContent([ 'status' => 'ERROR', 'messages' => $errors ]);
    }
    
    return $response;
});

// Deletes robots based on primary key
$app->delete('/api/robots/{id:[0-9]+}', function ($id) use ($app) {
    // Operation to delete a robot with id $id
    $phql = 'DELETE FROM Store\Toys\Robots WHERE id = :id:';
    $status = $app->modelsManager->executeQuery($phql, [ 'id' => $id ]);

    // Create a response
    $response = new Response();

    if ($status->success() === true) {
        $response->setJsonContent([ 'status' => 'OK' ]);
    } else {
        $errors = [];
        foreach ($status->getMessages() as $message) {
            $errors[] = $message->getMessage();
        }

        // Change the HTTP status
        $response->setStatusCode(409, 'Conflict');
        $response->setJsonContent([ 'status' => 'ERROR', 'messages' => $errors ]);
    }

    return $response;
});

$app->handle();
