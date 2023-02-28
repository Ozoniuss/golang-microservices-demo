Store
-----

This package contains the domain model, as well as packages to interact with various types of databases, such as an in-memory database or relational databases. It provides a `Store` interface which defines the possible operations on the database, and those methods must be implemented for each particular storage driver. This pattern with an interface describing the database operations reminds me a bit of the repository pattern that we've overused during college (especially in java classes), `Store` acting as an abstract repository whereas `inmemory` and `postgres` acting as concrete classes implementing the abstract repository. It's a pattern I didn't really see in production though (maybe because the operations were done on a single database and there was no need for it) but in this case, I felt that abstracting away the possible database operations would greatly simplify switching between databases at runtime through a single flag.

The `port.go` file contains a struct representing the port model. If there are multiple models, having a separate package with all the entities is a reasonable design choice, but in this case it doesn't really make a difference.

See:

- [inmemory](./inmemory/) for more information about the inmemory database code;
- [postgres](./postgres/) for more information about the postgres database code.

