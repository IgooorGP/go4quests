package entities

// Model of the transactions that users consumed points
type ConsumedPointsTransaction struct {
	BaseModel
	Points int    `gorm:"not null"` // saves points because rewards values may vary over time
	Reward Reward `gorm:"not null"` // one-to-one with a reward: rewards consume points
}
