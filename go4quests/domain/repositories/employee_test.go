package repositories

import (
	"github.com/IgooorGP/go4quests/go4quests/domain/entities"
	"github.com/IgooorGP/go4quests/go4quests/infra/persistence"
	"github.com/IgooorGP/go4quests/tests"
	"github.com/go-playground/assert/v2"
	"testing"
)

// Asserts
func TestSaveEmployee_Success(t *testing.T) {

	// test scenario
	db := tests.CreateTestDatabase()
	employee := &entities.Employee{
		FirstName:                  "i",
		LastName:                   "gp",
		Level:                      1,
		EarnedPointsTransactions:   nil,
		ConsumedPointsTransactions: nil,
	}

	// repository method invocation
	employeeRepository := persistence.NewEmployeeRepository(db)

	// assertions
	savedEmployee, _ := employeeRepository.SaveEmployee(employee)
	assert.Equal(t, employee, savedEmployee)
}
