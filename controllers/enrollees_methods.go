package controllers

import models "github.com/bhatia4/enrollee-services-go/models"
import "encoding/json" //to implement marshalling/unmarshalling of JSON data
import "net/http"
import "github.com/gin-gonic/gin"
import "github.com/satori/go.uuid"

func addEnrollee(c *gin.Context) {
	var inputEnrollee models.Enrollee
	c.BindJSON(&inputEnrollee);
	inputEnrollee.ID = uuid.NewV4()
	
	db := models.GetDBConn()
	db.Table("enrolleesdb.enrollees").Create(&inputEnrollee)

	outputEnrolleeBytes, _ := json.Marshal(inputEnrollee)
	c.JSON(http.StatusOK, string(outputEnrolleeBytes))
}

func updateEnrollee(c *gin.Context) {
	var inputEnrollee models.Enrollee
	var dbEnrollee models.Enrollee
	c.BindJSON(&inputEnrollee);

	enrolleeid, uuid_err := uuid.FromString( c.Param("enrolleeid") )
	if uuid_err != nil {
		c.JSON(http.StatusBadRequest, uuid_err)
		return
	}
	inputEnrollee.ID = enrolleeid
	
	db := models.GetDBConn()
	result := db.Table("enrolleesdb.enrollees").First(&dbEnrollee, inputEnrollee.ID) // find Enrollee with input enrolleeid (primary key)
	if (result.RowsAffected==1) { //if enrollee found using ID
		db.Table("enrolleesdb.enrollees").Save(inputEnrollee)
	}
	
	inputEnrolleeBytes, _ := json.Marshal(inputEnrollee)
	c.JSON(http.StatusOK, string(inputEnrolleeBytes))
}

func deleteEnrollee(c *gin.Context) {
	var dbEnrollee models.Enrollee
	
	enrolleeid, uuid_err := uuid.FromString( c.Param("enrolleeid") )
	if uuid_err != nil {
		c.JSON(http.StatusBadRequest, uuid_err)
		return
	}
	
	db := models.GetDBConn()
	result := db.Table("enrolleesdb.enrollees").First(&dbEnrollee, enrolleeid) // find Enrollee with input enrolleeid (primary key)
	if (result.RowsAffected==1) { //if enrollee found using ID
		db.Table("enrolleesdb.enrollees").Delete(&dbEnrollee, enrolleeid)
	}
	
	dbEnrolleeBytes, _ := json.Marshal(dbEnrollee)
	c.JSON(http.StatusOK, string(dbEnrolleeBytes))
}