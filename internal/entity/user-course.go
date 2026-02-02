package entity

import "time"

type UserCourse struct {
	ID             int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         int       `gorm:"index:idx_user_course,unique" json:"userId"`
	CourseID       int       `gorm:"index:idx_user_course,unique" json:"courseId"`
	Progress       int       `gorm:"default:0" json:"progress"`
	CurrentChapter int       `gorm:"default:1" json:"currentChapter"`
	IsCompleted    bool      `gorm:"default:false" json:"isCompleted"`
	TimeStarted    time.Time `gorm:"autoCreateTime" json:"timeStarted"`
	TimeFinished   time.Time `json:"timeFinished"`
	EnrolledAt     time.Time `gorm:"autoCreateTime" json:"enrolledAt"`

	Course Course `gorm:"foreignKey:CourseID" json:"course"`
}
