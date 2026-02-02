package entity

import "time"

type Course struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        string    `gorm:"not null" json:"code"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Image       string    `gorm:"default:''" json:"image"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	Chapters []Chapter `gorm:"foreignKey:CourseID" json:"chapters,omitempty"`
}