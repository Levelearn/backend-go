package entity

import "time"

type Trade struct {
	ID                int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title             string    `json:"title"`
	Image             string    `json:"image"`
	Description       string    `json:"description"`
	RequiredBadgeType BadgeType `gorm:"type:enum('BEGINNER','INTERMEDIATE','ADVANCE')" json:"requiredBadgeType"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
