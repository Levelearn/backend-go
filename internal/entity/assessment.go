package entity

import "time"

type Assessment struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ChapterID   int       `json:"chapterId"`
	Instruction string    `gorm:"type:text" json:"instruction"`
	Questions   []byte    `gorm:"type:json" json:"questions"` // Raw JSON
	Answers     []byte    `gorm:"type:json" json:"answers"`   // Raw JSON
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
