package entities

// Model of a quest that can bring points to employees
type Quest struct {
	BaseModel
	Description string `gorm:"size:255; not null;"`
	MinLevel    int    `gorm:"default: 1; not null"`
	Points      int    `gorm:"not null"`
}
