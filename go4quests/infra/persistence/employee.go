package persistence

import (
	"github.com/IgooorGP/go4quests/go4quests/domain/entities"
	"github.com/jinzhu/gorm"
)

type EmployeeRepository struct {
	database *gorm.DB
}

// Construction function
func NewEmployeeRepository(database *gorm.DB) *EmployeeRepository {
	repository := &EmployeeRepository{database: database}

	return repository
}

// Interface implementation
func (employeeRepository *EmployeeRepository) SaveEmployee(employee *entities.Employee) (*entities.Employee, error) {
	// transactions enabled by default: error is rollback'd
	if err := employeeRepository.database.Create(employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}
