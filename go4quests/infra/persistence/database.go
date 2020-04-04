package persistence

import (
	"fmt"
	"github.com/IgooorGP/go4quests/go4quests/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database abstraction with migration facility
type Database struct {
	conn *gorm.DB
}

func (database *Database) Migrate(model interface{}) error {
	err := database.conn.AutoMigrate(model).Error

	return err
}

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

// Wrapper on top of CreateDatabaseConnection that uses app context's env variables to open the connection
func CreateDatabaseConnectionWithAppContext() *Database {
	db := CreateDatabaseConnection(
		config.DatabaseEngine,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName,
		config.DatabaseAppUser,
		config.DatabaseAppPassword,
		config.DatabaseUseSSL,
	)

	databaseWrapper := &Database{conn: db}

	return databaseWrapper
}
