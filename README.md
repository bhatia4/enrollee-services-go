# Enrollees Services in GO

Coding completed by [Kunal Bhatia](https://github.com/bhatia4/), as as part of coding challenge (Challenge details [here](https://github.com/bhatia4/enrollee-services-go/blob/main/backend-challenge.md))

GO Libraries and modules used:
* [GORM](https://gorm.io/) - ORM library for Golang 
* [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) - GORM Postgres drivers
* [Gin Web Framework](https://github.com/gin-gonic/gin) - Helps to quickly build and scale up web services 
* [github.com/satori/go.uuid](https://github.com/satori/go.uuid) - UUID package lib

## Installation

Use the [build.bat](https://github.com/bhatia4/enrollee-services-go/blob/main/build.bat) or [build.sh](https://github.com/bhatia4/enrollee-services-go/blob/main/build.sh) command to setup and run the services

Using postman or any services client you can test against following endpoints:
* Add a new enrollee - HTTP POST localhost:8080/api/v1/enrollees/
* Modify an existing enrollee - HTTP PUT localhost:8080/api/v1/enrollees/:enrolleeid/
* Remove an enrollee entirely - HTTP DELETE localhost:8080/api/v1/enrollees/:enrolleeid/
* Add dependent to an enrollee - HTTP POST localhost:8080/api/v1/enrollees/:enrolleeid/dependent
* Modify existing dependent - HTTP PUT localhost:8080/api/v1/enrollees/:enrolleeid/dependent/:dependent
* Remove existing dependent from an enrollee - HTTP DELETE localhost:8080/api/v1/enrollees/:enrolleeid/dependent/:dependent

## Example JSON Request - Add a new enrollee / Modify an existing enrollee
```
{
  "activationStatus": true,
  "birthDate": "1987-02-29",
  "name": "Kenny Gee",
  "phoneNumber": "548-828-6089"
}
```

## Example JSON Request - Add dependent to an enrollee / Remove existing dependent from an enrollee
```
{
  "birthDate": "2020-09-10",
  "name": "John Smith"
}
```

## Requirements

Tested and run against Go >= 1.15.6. 
Uses Go modules

## Database Documentation
* Database model shown below:
<img src="https://raw.githubusercontent.com/bhatia4/enrollee-services-go/main/db/enrolleesdb%20DB%20model%20diagram.png">

* More details on tables above [here](https://htmlpreview.github.io/?https://github.com/bhatia4/enrollee-services-go/blob/master/db/enrolleesdb.html)
* DDL scripts on creating above tables on your own Postgres database server [here](https://github.com/bhatia4/enrollee-services-go/blob/master/db/enrolleesdb.sql)
* To update database server and client connection goto models/database.go file found [here](https://github.com/bhatia4/enrollee-services-go/blob/main/models/database.go). Appropriately change the constant string - dsn

## Links used as reference while coding
* [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
* [https://github.com/gin-gonic/gin#gin-web-framework](https://github.com/gin-gonic/gin#gin-web-framework)
* [https://godoc.org/github.com/gin-gonic/gin](https://godoc.org/github.com/gin-gonic/gin)
* [https://gobyexample.com/json](https://gobyexample.com/json)
* [https://github.com/satori/go.uuid](https://github.com/satori/go.uuid)
* [https://medium.com/@the.hasham.ali/how-to-use-uuid-key-type-with-gorm-cc00d4ec7100](https://medium.com/@the.hasham.ali/how-to-use-uuid-key-type-with-gorm-cc00d4ec7100)
* [https://golang.org/pkg/encoding/json/](https://golang.org/pkg/encoding/json/)
* [https://travix.io/type-embedding-in-go-ba40dd4264df](https://travix.io/type-embedding-in-go-ba40dd4264df)
* [https://daniel-dc.medium.com/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework-3d3f95ccf2e7](https://daniel-dc.medium.com/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework-3d3f95ccf2e7)
* [https://gorm.io/docs/connecting_to_the_database.html](https://gorm.io/docs/connecting_to_the_database.html)
* [https://gorm.io/docs/create.html](https://gorm.io/docs/create.html)
* [https://gorm.io/docs/models.html](https://gorm.io/docs/models.html)
* [https://gorm.io/docs/update.html](https://gorm.io/docs/update.html)
* [https://gorm.io/docs/conventions.html](https://gorm.io/docs/conventions.html)
* [https://gorm.io/docs/generic_interface.html](https://gorm.io/docs/generic_interface.html)
