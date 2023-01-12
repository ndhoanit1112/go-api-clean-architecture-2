# Go API Clean architecture
This repository contains an implementation of a Go application following Clean Architecture principles. The application is a simple API that demonstrates the use of Clean Architecture in Go.

## Directory structure
```
├── cmd
│   └── api
│       └── main.go
└── internal
    ├── adapters
    │   ├── http
    │   │   ├── handlers
    │   │   │   ├── handlers.go
    │   │   │   ├── product
    │   │   │   │   ├── handlers.go
    │   │   │   │   └── routes.go
    │   │   │   └── user
    │   │   └── server.go
    │   └── mysql
    │       ├── db.go
    │       └── repositories
    │           ├── product.go
    │           └── user.go
    ├── configs
    │   └── configs.go
    ├── entities
    │   ├── base.go
    │   └── product.go
    └── usecases
        ├── models
        │   └── product.go
        └── services
            └── product.go
```

I tend to follow the [Go standard project layout](https://github.com/golang-standards/project-layout), and try to implement the core idea of Clean architecture. The goal is to make the application's core business logic independent of any particular framework, database, or delivery mechanism.
This implementation is basically the same with [my previous implementation](https://github.com/ndhoanit1112/go-api-clean-architecture). I just try to change the name of application layers so that it closely relects the [Clean Architecture by Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). The Dependency rule is completely the same.

![Clean Architecture by Uncle Bob](https://github.com/mattia-battiston/clean-architecture-example/blob/master/docs/images/clean-architecture-diagram-2.png?raw=true)

[Image source](https://github.com/mattia-battiston/clean-architecture-example)

The main three layers are the Adapters, Use cases, and Entities.
### Adapters
The Domain layer contains the core business logic of the application. It defines the data structures and operations that represent the problem domain, as well as interfaces and services that define the contract between the domain and the other layers. This layer should be independent of any particular framework or delivery mechanism.

### Use cases
This layer provides an abstraction layer between the internal parts of the application and external services, such as a database, a cache, or an API. It provides a way to interact with the external service, without the internal parts of the application being directly aware of the implementation details. The adapters translate the data into a format that can be understood by the use cases and vice versa.

### Entities
This layer contains the structs or objects that represent the internal data models used in the application. The entities define the data structure and the business rules for the application. The entities should be independent of the other layers and should not contain any external dependencies. The entities are also independent of data access code, so it could be reused in different type of data storage.

Together, these layers form the core of the application. The adapters provide the communication to the external world, the use cases provide the business logic and rules, and the entities provide the data models. By separating these layers, the application becomes more flexible, maintainable, and testable. The outer layer can be easily replaced or updated, without affecting the inner layers, and the inner layers can be reused in different contexts.


## References
[The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

[Common web application architectures](https://learn.microsoft.com/en-us/dotnet/architecture/modern-web-apps-azure/common-web-application-architectures#clean-architecture)

[Clean Architecture Example](https://github.com/mattia-battiston/clean-architecture-example)

## Todo
- [ ] General response format
- [ ] Logging
- [ ] Error handling
- [ ] Adding Authentication
