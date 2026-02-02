package entity

import "time"

type Assignment struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ChapterID   int       `json:"chapterId"`
	Instruction string    `gorm:"type:longtext" json:"instruction"`
	FileURL     string    `json:"fileUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
