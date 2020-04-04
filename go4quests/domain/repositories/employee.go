package repositories

import (
	"github.com/IgooorGP/go4quests/go4quests/domain/entities"
	uuid "github.com/satori/go.uuid"
)

// Repository for Employee entity interaction with database
type EmployeeRepository interface {
	SaveEmployee(employee *entities.Employee) (*entities.Employee, error)
	GetEmployee(uuid uuid.UUID) (*entities.Employee, error)

	// TODO: implement later
	// GetAllEmployee() ([]entities.Employee, error)
	// UpdateEmployee(employee *entities.Employee) (*entities.Employee, error)
	// DeleteEmployee(uuid uuid.UUID) error
}
