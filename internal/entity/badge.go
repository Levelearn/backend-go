package entity

import "time"

type BadgeType string

const (
	BadgeBeginner     BadgeType = "BEGINNER"
	BadgeIntermediate BadgeType = "INTERMEDIATE"
	BadgeAdvance      BadgeType = "ADVANCE"
)

type Badge struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Type      BadgeType `gorm:"type:enum('BEGINNER','INTERMEDIATE','ADVANCE')" json:"type"`
	Image     string    `json:"image"`
	CourseID  int       `json:"courseId"`
	ChapterID int       `json:"chapterId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
