package persistence

import (
	"fmt"
	"github.com/IgooorGP/go4quests/go4quests/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database abstraction with migration facility
type Database struct {
	DatabaseConfig config.DatabaseConfig
	Conn           *gorm.DB
}

// Database interface creation
func NewDatabase(databaseConfig config.DatabaseConfig) *Database {
	database := &Database{DatabaseConfig: databaseConfig, Conn: nil}

	return database
}

// shortcut to connect to the database
func (database *Database) Connect() {
	databaseConfig := database.DatabaseConfig

	// Url connection string
	var DatabaseURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		databaseConfig.DatabaseHost,
		databaseConfig.DatabasePort,
		databaseConfig.DatabaseUser,
		databaseConfig.DatabaseName,
		databaseConfig.DatabasePassword,
		databaseConfig.DatabaseUseSSL,
	)

	// actual connection to the database server
	conn, err := gorm.Open(databaseConfig.DatabaseEngine, DatabaseURL)

	// unable to proceed if we cannot connect to the database
	if err != nil {
		panic(err)
	}

	// sets the database connection on the struct
	database.Conn = conn
}
