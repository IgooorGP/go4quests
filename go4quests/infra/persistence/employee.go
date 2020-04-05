package persistence

import (
	"github.com/IgooorGP/go4quests/go4quests/domain/entities"
)

type EmployeeRepository struct {
	database *Database
}

// Construction function
func NewEmployeeRepository(database *Database) *EmployeeRepository {
	repository := &EmployeeRepository{database: database}

	return repository
}

// Interface implementation
func (employeeRepository *EmployeeRepository) SaveEmployee(employee *entities.Employee) (*entities.Employee, error) {
	// transactions enabled by default: error is rollback'd
	if err := employeeRepository.database.Conn.Create(employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}
