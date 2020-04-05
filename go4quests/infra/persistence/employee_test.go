package persistence

import (
	"github.com/IgooorGP/go4quests/go4quests/config"
	"github.com/IgooorGP/go4quests/go4quests/domain/entities"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSaveEmployee_ShouldSuccessfullyCreateEmployee(t *testing.T) {

	// test scenario: relation creation + seeds + deferred conn close
	db := NewDatabase(config.DBConfig)
	db.Connect()
	db.Conn.AutoMigrate(&entities.Employee{})
	defer db.Disconnect()

	employee := &entities.Employee{
		FirstName:                  "i",
		LastName:                   "gp",
		Level:                      1,
		EarnedPointsTransactions:   nil,
		ConsumedPointsTransactions: nil,
	}

	// repository method invocation
	employeeRepository := NewEmployeeRepository(db)

	// assertions
	savedEmployee, _ := employeeRepository.SaveEmployee(employee)
	assert.Equal(t, employee, savedEmployee)
}

func TestSaveEmployee_ShouldSuccessfullyGetEmployee(t *testing.T) {

	// test scenario: previously created employee
	db := NewDatabase(config.DBConfig)
	db.Connect()
	db.Conn.AutoMigrate(&entities.Employee{})
	defer db.Disconnect()

	// seed for the database scenario
	employee := &entities.Employee{
		FirstName:                  "i",
		LastName:                   "gp",
		Level:                      1,
		EarnedPointsTransactions:   nil,
		ConsumedPointsTransactions: nil,
	}

	if err := db.Conn.Create(employee).Error; err != nil {
		panic(err)
	}

	employeeID := employee.ID

	// function invocation
	employeeRepository := NewEmployeeRepository(db)
	employeeFromDatabase, err := employeeRepository.GetEmployeeByID(employeeID)

	// assertions
	assert.Equal(t, err, nil)
	assert.Equal(t, employeeFromDatabase.ID, employee.ID)
	assert.Equal(t, employeeFromDatabase.FirstName, employee.FirstName)
	assert.Equal(t, employeeFromDatabase.LastName, employee.LastName)
}

func TestSaveEmployee_ShouldSuccessfullyNotFindEmployee(t *testing.T) {

	// test scenario: previously created employee
	db := NewDatabase(config.DBConfig)
	db.Connect()
	db.Conn.AutoMigrate(&entities.Employee{})
	defer db.Disconnect()

	// seed for the database scenario
	employee := &entities.Employee{
		FirstName:                  "i",
		LastName:                   "gp",
		Level:                      1,
		EarnedPointsTransactions:   nil,
		ConsumedPointsTransactions: nil,
	}

	if err := db.Conn.Create(employee).Error; err != nil {
		panic(err)
	}

	// function invocation
	employeeRepository := NewEmployeeRepository(db)
	employeeFromDatabase, err := employeeRepository.GetEmployeeByID(0) // unkown id

	// assertions
	assert.Equal(t, err.Error(), "record not found")
	assert.Equal(t, employeeFromDatabase, nil)
}
