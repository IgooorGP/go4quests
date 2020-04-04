package entities

import "time"

// Base model for entities with UUID id, created_at, updated_at and deleted_at
type BaseModel struct {
	ID        uint64    `gorm:"primary_key; auto_increment"`
	CreatedAt time.Time `gorm:"index; default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `gorm:"index; default:CURRENT_TIMESTAMP;"`
	DeletedAt *time.Time
}
