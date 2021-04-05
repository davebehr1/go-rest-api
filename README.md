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