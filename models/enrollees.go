package models

import "github.com/satori/go.uuid"

type Applicant struct {
    Name				string `json:"name"`
    BirthDate			string `json:"birthDate" gorm:"column:birthdate"`
}

type Enrollee struct {
    Applicant
    ID					uuid.UUID `json:"id" gorm:"primaryKey;column:enrolleeid"`
	ActivationStatus 	bool `json:"activationStatus" gorm:"column:activated"`
	PhoneNumber 		string `json:"phoneNumber" gorm:"column:phonenumber"`
}

type Dependent struct {
    Applicant
    ID					uuid.UUID `json:"id" gorm:"primaryKey;column:dependentid"`
	EnrolleeID			uuid.UUID `json:"enrolleeid" gorm:"column:enrolleeid"`
}