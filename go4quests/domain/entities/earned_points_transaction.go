package entities

// Model of the transactions that generated positive points
type EarnedPointsTransaction struct {
	BaseModel
	Points int `gorm:"not null"` // saves points because quests values may vary over time
	Quest  int `gorm:"not null"` // one-to-one with a quest: points always come from a quest
}
