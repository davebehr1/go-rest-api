# lxdAssessment

## BUILD INSTRUCTIONS:

* cd server:
   -  docker-compose up
   -  go run main.go migrate
   -  go run main.go serve
* cd client:
   - execute any of the cli commands listed below


## Cli commands for Server:

*   Serve: serves the books api over http
*   Migrate: creates database tables

## Cli commands for client:



*   ### Create:
    *   Collection:
        *   Flags:
            *   --collection / required
            *   --help / help for command
        *   Example:
            *   Go run main.go create collection --collection=’{“name”:”Fiction”}’
    *   Book
        *   Flags:
            *   --book/ required
            *   --collection / optional
            *   --help / help for command
        *   Example:
            *   Go run main.go create book --book=’{“title”:”harryPotter”,”author”:”J.K Rowling”, “description”:”book about magic”}’ --collection=fiction
*   ### Update:
    *   Collection
    *   Book
*   ### Help
*   ### List:
    *   Books:
        *   Flags:
            *   --author / optional
            *   --title / optional
            *   --fromDate / optional
            *   --toDate  / optional
        *   Example:
            *   Go run main.go list books --author=David --title=harry Potter
    *   Collections:
        *   Flags:
            *   --name
        *   Examples:
            *   Go run main.go list collections --name=Fiction
    *   --help

## REST API:
 
### RETURN VALUES:
   - success
   - error
#### success return:

```js
{
    "code": 200, // success codes
    "payload":{}
    "success":true
}
```

#### error:

```js
{
    "code": 401, // error codes
    "errorMessage": {}  // the error that occured, the client does not have to know about alot of these for security
    "success":false
}
```


### ROUTES:
   - [`/1.0`](###1.0):
      * [`/1.0/test`](###test):
      * [`/1.0/books`](###books):
      * [`/1.0/book`](###book): 
      * [`/1.0/collection`](###collection):
      * [`/1.0/collections`](###collections):


### `/1.0`
#### Version 1 of the api

### `/1.0/test`
#### GET
 * Description: test if the api is alive
 * Return: a string to say if the api is up and running

### `/1.0/books`
#### GET
 * Description: books in the database
 * Return: list of books


### `/1.0/book`
#### POST
 * Description: adds book to database
 * Return: returns book that was added to database

#### DELETE
 * Description: deletes a book from database
 * Return: returns a success message if deleted else the error is returned

#### PATCH
 * Description: update a book entity
 * Return: returns the book that was updated


### `/1.0/collections`
#### GET
 * Description: collections in the database
 * Return: list of collections

### `/1.0/collection`
#### PUT
 * Description: adds collection to database
 * Return: collection that was added to the database

#### DELETE
 * Description: deletes a collection from database
 * Return: returns a success message if deleted else the error is returned

#### PATCH
 * Description: update a collection entity
 * Return: returns the collection that was updated


 ## Database Structure:

### Book
| Field       	   | data type 	|             	|
|----------------	|-------------	|-------------	|
| id          	   | int8      	| FOREIGN_KEY 	|
| author      	   | varchar   	|             	|
| description 	   | varchar   	|             	|
| title       	   | varchar   	|             	|
| publishedAt   	| timestamp 	|             	|
| updatedAt   	   | timestamp 	|             	|

### Collection
| Field     	| data type 	|             	|
|-----------	|-----------	|-------------	|
| id        	| int8      	| PRIMARY_KEY 	|
| name      	| varchar   	|             	|
| publishedAt 	| timestamp 	|             	|
| updatedAt 	| timestamp 	|             	|
