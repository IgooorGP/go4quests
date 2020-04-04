package persistence

import (
	"github.com/IgooorGP/go4quests/go4quests/domain/entities"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSaveEmployee_ShouldSuccessfullyCreateEmployee(t *testing.T) {

	// test scenario: relation creation + seeds + deferred conn close
	db := CreateDatabaseConnectionWithAppContext()
	db.conn.AutoMigrate(entities.Employee{})
	defer db.conn.Close()

	employee := &entities.Employee{
		FirstName:                  "i",
		LastName:                   "gp",
		Level:                      1,
		EarnedPointsTransactions:   nil,
		ConsumedPointsTransactions: nil,
	}

	// repository method invocation
	employeeRepository := NewEmployeeRepository(db.conn)

	// assertions
	savedEmployee, _ := employeeRepository.SaveEmployee(employee)
	assert.Equal(t, employee, savedEmployee)
}
