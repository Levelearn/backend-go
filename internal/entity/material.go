package entity

import "time"

type Material struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ChapterID int       `json:"chapterId"`
	Name      string    `json:"name"`
	Content   string    `gorm:"type:longtext" json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
