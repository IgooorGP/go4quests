package entities

// Model of an employee
type Employee struct {
	BaseModel
	FirstName                  string                       `gorm:"size:150; not null;"`
	LastName                   string                       `gorm:"size:150; not null;"`
	Level                      int                          `gorm:"default:1; not null;"` // level one as the initial one
	EarnedPointsTransactions   []EarnedPointsTransactions   // one-to-many positive transactions
	ConsumedPointsTransactions []ConsumedPointsTransactions // one-to-many negative transactions
}

// Model of a quest that can bring points to employees
type Quest struct {
	BaseModel
	Description string `gorm:"size:255; not null;"`
	MinLevel    int    `gorm:"default: 1; not null"`
	Points      int    `gorm:"not null"`
}

// Model of a quest that can bring points to employees
type Reward struct {
	BaseModel
	Description string `gorm:"size:255; not null;"`
	MinLevel    int    `gorm:"default: 1; not null"`
	Points      int    `gorm:"not null"`
}

// Model of the transactions that generated positive points
type EarnedPointsTransactions struct {
	BaseModel
	Points int `gorm:"not null"` // saves points because quests values may vary over time
	Quest  int `gorm:"not null"` // one-to-one with a quest: points always come from a quest
}

// Model of the transactions that users consumed points
type ConsumedPointsTransactions struct {
	BaseModel
	Points int    `gorm:"not null"` // saves points because rewards values may vary over time
	Reward Reward `gorm:"not null"` // one-to-one with a reward: rewards consume points
}
