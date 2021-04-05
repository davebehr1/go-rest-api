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