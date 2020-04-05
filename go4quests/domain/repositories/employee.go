package repositories

import (
	"github.com/IgooorGP/go4quests/go4quests/domain/entities"
)

// Repository for Employee entity interaction with database
type IEmployeeRepository interface {
	SaveEmployee(employee *entities.Employee) (*entities.Employee, error)
	GetEmployeeByID(id int) (*entities.Employee, error)

	// TODO: implement later
	// GetAllEmployee() ([]entities.Employee, error)
	// UpdateEmployee(employee *entities.Employee) (*entities.Employee, error)
	// DeleteEmployee(uuid uuid.UUID) error
}
