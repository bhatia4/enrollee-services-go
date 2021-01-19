# Enrollees Services in GO

Coding completed by [Kunal Bhatia](https://github.com/bhatia4/), as as part of coding challenge (Challenge details [here](https://github.com/bhatia4/enrollee-services-go/blob/main/backend-challenge.md))

GO Libraries and modules used:
* [GORM](https://gorm.io/) - ORM library for Golang 
* [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) - GORM Postgres drivers
* [Gin Web Framework](github.com/gin-gonic/gin) - Helps to quickly build and scale up web services 
* [github.com/satori/go.uuid](github.com/satori/go.uuid) - UUID package lib

## Installation

Use the [build.bat](https://github.com/bhatia4/enrollee-services-go/blob/main/build.bat) or [build.sh](https://github.com/bhatia4/enrollee-services-go/blob/main/build.sh) command to setup and run the services

Using postman or any services client you can test against following endpoints:
* Add a new enrollee - HTTP POST localhost:8080/api/v1/enrollees/
* Modify an existing enrollee - HTTP PUT localhost:8080/api/v1/enrollees/:enrolleeid/
* Remove an enrollee entirely - HTTP DELETE localhost:8080/api/v1/enrollees/:enrolleeid/
* Add dependents to an enrollee - HTTP POST localhost:8080/api/v1/enrollees/:enrolleeid/dependent
* Modify existing dependents - HTTP PUT localhost:8080/api/v1/enrollees/:enrolleeid/dependent/:dependent
* Add dependents to an enrollee - HTTP DELETE localhost:8080/api/v1/enrollees/:enrolleeid/dependent/:dependent

## Example JSON Request - Add a new enrollee / Modify an existing enrollee
{
  "activationStatus": true,
  "birthDate": "1987-02-29",
  "name": "Kenny Gee",
  "phoneNumber": "548-828-6089"
}

## Requirements

Tested and run against Go >= 1.15.6. 
Uses Go modules

## Links
* [RFC 4122](http://tools.ietf.org/html/rfc4122)
* [DCE 1.1: Authentication and Security Services](http://pubs.opengroup.org/onlinepubs/9696989899/chap5.htm#tagcjh_08_02_01_01)


* [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
* [https://github.com/gin-gonic/gin#gin-web-framework](https://github.com/gin-gonic/gin#gin-web-framework)
* [https://godoc.org/github.com/gin-gonic/gin](https://godoc.org/github.com/gin-gonic/gin)
* [https://gobyexample.com/json](https://gobyexample.com/json)
* [https://github.com/satori/go.uuid](https://github.com/satori/go.uuid)
* [https://medium.com/@the.hasham.ali/how-to-use-uuid-key-type-with-gorm-cc00d4ec7100](https://medium.com/@the.hasham.ali/how-to-use-uuid-key-type-with-gorm-cc00d4ec7100)
* [https://golang.org/pkg/encoding/json/](https://golang.org/pkg/encoding/json/)
* 
* [https://travix.io/type-embedding-in-go-ba40dd4264df](https://travix.io/type-embedding-in-go-ba40dd4264df)
* [https://daniel-dc.medium.com/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework-3d3f95ccf2e7](https://daniel-dc.medium.com/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework-3d3f95ccf2e7)
* [https://gorm.io/docs/connecting_to_the_database.html](https://gorm.io/docs/connecting_to_the_database.html)
* [https://gorm.io/docs/create.html](https://gorm.io/docs/create.html)
* [https://gorm.io/docs/models.html](https://gorm.io/docs/models.html)
* [https://gorm.io/docs/update.html](https://gorm.io/docs/update.html)
* [https://gorm.io/docs/conventions.html](https://gorm.io/docs/conventions.html)
* [https://gorm.io/docs/generic_interface.html](https://gorm.io/docs/generic_interface.html)