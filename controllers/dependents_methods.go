package controllers

import models "github.com/bhatia4/enrollee-services-go/models"
import "encoding/json" //to implement marshalling/unmarshalling of JSON data
import "net/http"
import "github.com/gin-gonic/gin"
import "github.com/satori/go.uuid"

func addDependent(c *gin.Context) {
	var inputDependent models.Dependent
	var dbEnrollee models.Enrollee
	c.BindJSON(&inputDependent);
	inputDependent.ID = uuid.NewV4()

	enrolleeid, uuid_err := uuid.FromString( c.Param("enrolleeid") )
	if uuid_err != nil {
		c.JSON(http.StatusBadRequest, uuid_err)
		return
	}
	inputDependent.EnrolleeID = enrolleeid

	db := models.GetDBConn()
	result := db.Table("enrolleesdb.enrollees").First(&dbEnrollee, inputDependent.EnrolleeID) // find Enrollee with input enrolleeid (foreign key)
	if (result.RowsAffected==1) { //if enrollee found using enrolleeid
		db.Table("enrolleesdb.dependents").Create(&inputDependent)
	}
	
	outputDependentBytes, _ := json.Marshal(inputDependent)
	c.JSON(http.StatusOK, string(outputDependentBytes))
}

func updateDependent(c *gin.Context) {
	var inputDependent models.Dependent
	var dbEnrollee models.Enrollee
	var dbDependent models.Dependent
	c.BindJSON(&inputDependent);

	dependentid, uuid_err1 := uuid.FromString( c.Param("dependentid") )
	if uuid_err1 != nil {
		c.JSON(http.StatusBadRequest, uuid_err1)
		return
	}
	inputDependent.ID = dependentid
	
	enrolleeid, uuid_err2 := uuid.FromString( c.Param("enrolleeid") )
	if uuid_err2 != nil {
		c.JSON(http.StatusBadRequest, uuid_err2)
		return
	}
	inputDependent.EnrolleeID = enrolleeid
	
	db := models.GetDBConn()
	result1 := db.Table("enrolleesdb.enrollees").First(&dbEnrollee, inputDependent.EnrolleeID) // find Enrollee with input enrolleeid (foreign key)
	if (result1.RowsAffected==1) { //if enrollee found using enrolleeid
	
		result2 := db.Table("enrolleesdb.dependents").First(&dbDependent, inputDependent.ID) // find Dependent with input id (primary key)
		if (result2.RowsAffected==1) { //if dependent found using dependentid {
			db.Table("enrolleesdb.dependents").Save(inputDependent)
		}
	}
	
	inputDependentBytes, _ := json.Marshal(inputDependent)
	c.JSON(http.StatusOK, string(inputDependentBytes))
}

func deleteDependent(c *gin.Context) {
	var dbEnrollee models.Enrollee
	var dbDependent models.Dependent
	
	dependentid, uuid_err1 := uuid.FromString( c.Param("dependentid") )
	if uuid_err1 != nil {
		c.JSON(http.StatusBadRequest, uuid_err1)
		return
	}

	enrolleeid, uuid_err2 := uuid.FromString( c.Param("enrolleeid") )
	if uuid_err2 != nil {
		c.JSON(http.StatusBadRequest, uuid_err2)
		return
	}
	
	db := models.GetDBConn()
	result1 := db.Table("enrolleesdb.enrollees").First(&dbEnrollee, enrolleeid) // find Enrollee with input enrolleeid (foreign key)
	if (result1.RowsAffected==1) { //if enrollee found using enrolleeid
	
		result2 := db.Table("enrolleesdb.dependents").First(&dbDependent, dependentid) // find Dependent with input id (primary key)
		if (result2.RowsAffected==1) { //if dependent found using dependentid {
			db.Table("enrolleesdb.dependents").Delete(&dbDependent, dependentid)
		}
	}
	
	dbEnrolleeBytes, _ := json.Marshal(dbDependent)
	c.JSON(http.StatusOK, string(dbEnrolleeBytes))
}