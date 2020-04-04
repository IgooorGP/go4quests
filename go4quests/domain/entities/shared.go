package entities

import "time"

// Base model for entities with UUID id, created_at, updated_at and deleted_at
type BaseModel struct {
	ID        string    `sql:"type:uuid; primary_key; default:uuid_generate_v4();"`
	CreatedAt time.Time `gorm:"index; default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `gorm:"index; default:CURRENT_TIMESTAMP;"`
	DeletedAt *time.Time
}
