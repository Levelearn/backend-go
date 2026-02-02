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

	Material   *Material   `gorm:"constraint:OnUpdate:CASCADE" json:"material,omitempty"`
	Assessment *Assessment `gorm:"constraint:OnUpdate:CASCADE" json:"assessment,omitempty"`
	Assignment *Assignment `gorm:"constraint:OnUpdate:CASCADE" json:"assignment,omitempty"`
}
