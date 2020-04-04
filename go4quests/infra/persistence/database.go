package persistence

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func CreateDatabaseConnection(
	engine string, host string, port string,
	dbName string, user string, password string, sslMode string,
) *gorm.DB {
	// Url connection string
	var DatabaseURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host,
		port,
		user,
		dbName,
		password,
		sslMode,
	)

	db, err := gorm.Open(engine, DatabaseURL)

	// unable to proceed if we cannot connect to the database
	if err != nil {
		panic(err)
	}

	return db
}
