package entities

// Model of an employee
type Employee struct {
	BaseModel
	FirstName                  string                      `gorm:"size:150; not null;"`
	LastName                   string                      `gorm:"size:150; not null;"`
	Level                      int                         `gorm:"default:1; not null;"` // level one as the initial one
	EarnedPointsTransactions   []EarnedPointsTransaction   // one-to-many positive transactions
	ConsumedPointsTransactions []ConsumedPointsTransaction // one-to-many negative transactions
}
