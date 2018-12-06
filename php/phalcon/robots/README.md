# Tutorial: Creating a Simple REST API

Reference: 
1. https://docs.phalconphp.com/en/3.3/tutorial-rest
2. bug fix: 
https://forum.phalconphp.com/discussion/15759/inclusionin-and-uniqueness-validator-errors
https://forum.phalconphp.com/discussion/8793/generated-orm-and-validation-is-not-compatible-to-itself


In this tutorial, we will explain how to create a simple application that provides a RESTful API using the different HTTP methods:

GET to retrieve and search data
POST to add data
PUT to update data
DELETE to delete data

==================================================================================
Method	URL							Action
----------------------------------------------------------------------------------
GET		/api/robots					Retrieves all robots
GET		/api/robots/search/Astro	Searches for robots with 'Astro' in their name
GET		/api/robots/2				Retrieves robots based on primary key
POST	/api/robots					Adds a new robot
PUT		/api/robots/2				Updates robots based on primary key
DELETE	/api/robots/2				Deletes robots based on primary key
==================================================================================


## To Clone This Repo
1. go to the projects dir in ur pc,
2. run the following command in cmd line:
git clone https://hecode@bitbucket.org/hecode/phalcon-rest-api-tutorial.git


## Testing our Application

Testing our Application on GitHub
Using curl we'll test every route in our application verifying its proper operation.


### Insert several new robots:
curl -i -X POST -d '{"name":"Robotina", "type":"droid", "year":1977}' http://localhost/phalcon-rest-api-tutorial/api/robots
curl -i -X POST -d '{"name":"Astro Boy", "type":"droid", "year":1977}' http://localhost/phalcon-rest-api-tutorial/api/robots
curl -i -X POST -d '{"name":"Terminator", "type":"droid", "year":1977}' http://localhost/phalcon-rest-api-tutorial/api/robots

HTTP/1.1 201 Created
Date: Tue, 21 Jul 2015 07:15:09 GMT
Server: Apache/2.2.22 (Unix) DAV/2
Content-Length: 75
Content-Type: text/html; charset=UTF-8

{"status":"OK","data":{"name":"C-3PO","type":"droid","year":1977,"id":"4"}}


### Obtain all the robots:

curl -i -X GET http://localhost/phalcon-rest-api-tutorial/api/robots

HTTP/1.1 200 OK
Date: Tue, 21 Jul 2015 07:05:13 GMT
Server: Apache/2.2.22 (Unix) DAV/2
Content-Length: 117
Content-Type: text/html; charset=UTF-8

[{"id":"1","name":"Robotina"},{"id":"2","name":"Astro Boy"},{"id":"3","name":"Terminator"}]

### Search a robot by its name:

curl -i -X GET http://localhost/phalcon-rest-api-tutorial/api/robots/search/Astro

HTTP/1.1 200 OK
Date: Tue, 21 Jul 2015 07:09:23 GMT
Server: Apache/2.2.22 (Unix) DAV/2
Content-Length: 31
Content-Type: text/html; charset=UTF-8

[{"id":"2","name":"Astro Boy"}]


### Obtain a robot by its id:

curl -i -X GET http://localhost/phalcon-rest-api-tutorial/api/robots/3

HTTP/1.1 200 OK
Date: Tue, 21 Jul 2015 07:12:18 GMT
Server: Apache/2.2.22 (Unix) DAV/2
Content-Length: 56
Content-Type: text/html; charset=UTF-8

{"status":"FOUND","data":{"id":"3","name":"Terminator"}}


### Try to insert a new robot:

curl -i -X POST -d '{"name":"C-3PO","type":"droid","year":1977}' http://localhost/phalcon-rest-api-tutorial/api/robots

HTTP/1.1 201 Created
Date: Tue, 20 Feb 2018 17:19:42 GMT
Server: Apache/2.4.25 (Win32) OpenSSL/1.0.2j PHP/7.1.4
X-Powered-By: PHP/7.1.4
Status: 201 Created
Content-Length: 75
Content-Type: application/json; charset=UTF-8

{"status":"OK","data":{"name":"C-3PO","type":"droid","year":1977,"id":"5"}}


### Try to insert a new robot with the name of an existing robot:

curl -i -X POST -d '{"name":"C-3PO","type":"droid","year":1977}' http://localhost/phalcon-rest-api-tutorial/api/robots

HTTP/1.1 409 Conflict
Date: Tue, 21 Jul 2015 07:18:28 GMT
Server: Apache/2.2.22 (Unix) DAV/2
Content-Length: 63
Content-Type: text/html; charset=UTF-8

{"status":"ERROR","messages":["The robot name must be unique"]}


### Or update a robot with an unknown type:

curl -i -X PUT -d '{"name":"ASIMO","type":"humanoid","year":2000}' http://localhost/phalcon-rest-api-tutorial/api/robots/4

HTTP/1.1 409 Conflict
Date: Tue, 21 Jul 2015 08:48:01 GMT
Server: Apache/2.2.22 (Unix) DAV/2
Content-Length: 104
Content-Type: text/html; charset=UTF-8

{"status":"ERROR","messages":["Value of field 'type' must be part of
    list: droid, mechanical, virtual"]}


### Finally, delete a robot:

curl -i -X DELETE http://localhost/phalcon-rest-api-tutorial/api/robots/4

HTTP/1.1 200 OK
Date: Tue, 21 Jul 2015 08:49:29 GMT
Server: Apache/2.2.22 (Unix) DAV/2
Content-Length: 15
Content-Type: text/html; charset=UTF-8

{"status":"OK"}