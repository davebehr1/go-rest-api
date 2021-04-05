# lxdAssessment


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
   - standard return
   - error
#### standard return:

```js
{
    "status_code": 200, // success codes
    "payload":{}
}
```

#### error:

```js
{
    "status_code": 401, // error codes
    "error": {}  // the error that occured, the client does not have to know about alot of these for security
}
```


### ROUTES:
   - [`/1.0`](###1.0):
      * [`/1.0/books`](###books):
      * [`/1.0/book/{id}`](###booksid):
      * [`/1.0/book`](###book): 
      * [`/1.0/collection`](###collection):
      * [`/1.0/collections`](###collections):
      * [`/1.0/collection/{id}`](###collectionsId):


### `/1.0`
#### Version 1 of the api

### `/1.0/books`
#### GET
 * Description: books in the database
 * Return: list of books


### `/1.0/book`
#### POST
 * Description: adds book to database
 * Return: returns book that was added to database


### `/1.0/book/{id}`
#### GET
 * Description: book with specified id
 * Return: book with specified id

### `/1.0/collections`
#### GET
 * Description: collections in the database
 * Return: list of collections

### `/1.0/collection`
#### PUT
 * Description: adds collection to database
 * Return: collection that was added to the database

### `/1.0/collection/{id}`
#### GET
 * Description: collection with specific id and other query parameters
 * Return: a collection with the id and other query parameters


 ## Database Structure:

### Book
| Field       	| data type 	|             	|
|-------------	|-----------	|-------------	|
| id          	| int8      	| FOREIGN_KEY 	|
| author      	| varchar   	|             	|
| description 	| varchar   	|             	|
| title       	| varchar   	|             	|
| createdAt   	| timestamp 	|             	|
| updatedAt   	| timestamp 	|             	|

### Collection
| Field     	| data type 	|             	|
|-----------	|-----------	|-------------	|
| id        	| int8      	| PRIMARY_KEY 	|
| name      	| varchar   	|             	|
| createdAt 	| timestamp 	|             	|
| updatedAt 	| timestamp 	|             	|
