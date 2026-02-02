package entity

import "time"

type Chapter struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Description  string    `json:"description"`
	Level        int       `json:"level"`
	CourseID     int       `json:"courseId"`
	IsCheckpoint int       `gorm:"default:0" json:"isCheckpoint"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`

	Materials   []Material   `gorm:"foreignKey:ChapterID" json:"materials,omitempty"`
	Assessments []Assessment `gorm:"foreignKey:ChapterID" json:"assessments,omitempty"`
	Assignments []Assignment `gorm:"foreignKey:ChapterID" json:"assignments,omitempty"`
}
