package entity

import "time"

type UserChapter struct {
	ID               int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           int       `json:"userId"`
	ChapterID        int       `json:"chapterId"`
	IsStarted        bool      `gorm:"default:false" json:"isStarted"`
	IsCompleted      bool      `gorm:"default:false" json:"isCompleted"`
	MaterialDone     bool      `gorm:"default:false" json:"materialDone"`
	AssessmentDone   bool      `gorm:"default:false" json:"assessmentDone"`
	AssignmentDone   bool      `gorm:"default:false" json:"assignmentDone"`
	AssessmentAnswer []byte    `gorm:"type:json" json:"assessmentAnswer"`
	AssessmentGrade  int       `gorm:"default:0" json:"assessmentGrade"`
	Submission       string    `json:"submission"`
	TimeStarted      time.Time `json:"timeStarted"`
	TimeFinished     time.Time `json:"timeFinished"`
	AssignmentScore  int       `gorm:"default:0" json:"assignmentScore"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
