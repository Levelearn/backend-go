package entity

import (
	"time"
)

// Enums kita handle sebagai string constant atau type validation di service level
type Role string
const (
	RoleStudent    Role = "STUDENT"
	RoleInstructor Role = "INSTRUCTOR"
	RoleAdmin      Role = "ADMIN"
)

type BadgeType string
const (
	BadgeBeginner     BadgeType = "BEGINNER"
	BadgeIntermediate BadgeType = "INTERMEDIATE"
	BadgeAdvance      BadgeType = "ADVANCE"
)

// --- USER ---
type User struct {
	ID                int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username          string    `gorm:"unique;not null" json:"username"`
	Password          string    `gorm:"not null" json:"-"` // Hide password from JSON
	Name              string    `gorm:"not null" json:"name"`
	Role              Role      `gorm:"type:enum('STUDENT','INSTRUCTOR','ADMIN');default:'STUDENT'" json:"role"`
	StudentID         *string   `json:"studentId"` // Pointer allow null
	Points            int       `gorm:"default:0" json:"points"`
	TotalCourses      int       `gorm:"default:0" json:"totalCourses"`
	BadgesCount       int       `gorm:"column:badges;default:0" json:"badgesCount"` // Renamed field to avoid clash with Relation
	InstructorID      *string   `json:"instructorId"`
	InstructorCourses int       `json:"instructorCourses"`
	Image             string    `gorm:"default:''" json:"image"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`

	// Relations
	EnrolledCourses []UserCourse  `gorm:"foreignKey:UserID" json:"enrolledCourses,omitempty"`
	UserBadges      []UserBadge   `gorm:"foreignKey:UserID" json:"userBadges,omitempty"`
}

// --- COURSE DOMAIN ---
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

type Material struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ChapterID int       `json:"chapterId"`
	Name      string    `json:"name"`
	Content   string    `gorm:"type:longtext" json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Menggunakan tipe JSON untuk Question/Answers
type Assessment struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ChapterID   int       `json:"chapterId"`
	Instruction string    `gorm:"type:text" json:"instruction"`
	Questions   []byte    `gorm:"type:json" json:"questions"` // Raw JSON
	Answers     []byte    `gorm:"type:json" json:"answers"`   // Raw JSON
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Assignment struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ChapterID   int       `json:"chapterId"`
	Instruction string    `gorm:"type:longtext" json:"instruction"`
	FileURL     string    `json:"fileUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// --- PROGRESS / GAMIFICATION ---
type UserCourse struct {
	ID             int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         int       `gorm:"index:idx_user_course,unique" json:"userId"`
	CourseID       int       `gorm:"index:idx_user_course,unique" json:"courseId"`
	Progress       int       `gorm:"default:0" json:"progress"`
	CurrentChapter int       `gorm:"default:1" json:"currentChapter"`
	IsCompleted    bool      `gorm:"default:false" json:"isCompleted"`
	TimeStarted    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"timeStarted"`
	TimeFinished   time.Time `json:"timeFinished"`
	EnrolledAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"enrolledAt"`

	Course Course `gorm:"foreignKey:CourseID" json:"course"`
}

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

type UserBadge struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int       `json:"userId"`
	BadgeID     int       `json:"badgeId"`
	IsPurchased bool      `gorm:"default:false" json:"isPurchased"`
	AwardedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"awardedAt"`

	Badge Badge `gorm:"foreignKey:BadgeID" json:"badge"`
}

type Trade struct {
	ID                int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title             string    `json:"title"`
	Image             string    `json:"image"`
	Description       string    `json:"description"`
	RequiredBadgeType BadgeType `gorm:"type:enum('BEGINNER','INTERMEDIATE','ADVANCE')" json:"requiredBadgeType"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type UserTrade struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int       `json:"userId"`
	TradeID     int       `json:"tradeId"`
	PurchasedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"purchasedAt"`
}