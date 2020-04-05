package persistence

import (
	"github.com/IgooorGP/go4quests/go4quests/domain/entities"
	"github.com/jinzhu/gorm"
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

func (employeeRepository *EmployeeRepository) GetEmployeeByID(id uint64) (*entities.Employee, error) {
	sqlQueryWhere := "id = ?"
	employee := &entities.Employee{} // empty struct to be populated after GET by Id

	err := employeeRepository.database.Conn.Where(sqlQueryWhere, id).First(employee).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// maybe employee was not found by that ID
			return nil, err
		} else {
			// unexpected error
			panic(err)
		}
	}

	return employee, nil
}
