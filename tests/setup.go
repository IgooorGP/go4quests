package tests

import (
	"github.com/IgooorGP/go4quests/go4quests/config"
	"github.com/IgooorGP/go4quests/go4quests/domain/entities"
	"github.com/IgooorGP/go4quests/go4quests/infra/persistence"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Creates a database connection AND migrates the models
func CreateTestDatabase() *gorm.DB {

	db := persistence.CreateDatabaseConnection(
		config.DatabaseEngine,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName,
		config.DatabaseAppUser,
		config.DatabaseAppPassword,
		config.DatabaseUseSSL,
	)

	if migrationError := db.AutoMigrate(entities.Employee{}).Error; migrationError != nil {
		panic(migrationError)
	}

	return db
}
