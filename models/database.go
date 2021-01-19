package models

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

const dsn string = "host=198.211.96.193 user=enrolleesdbusr1 password=12345 dbname=enrolleesdb port=5432 sslmode=disable"

func GetDBConn() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}