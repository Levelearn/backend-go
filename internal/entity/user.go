package entity

import "time"

type User struct {
	ID                int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username          string    `gorm:"unique;not null" json:"username"`
	Password          string    `gorm:"not null" json:"-"` // Hide password from JSON
	Name              string    `gorm:"not null" json:"name"`
	Role              Role      `gorm:"type:enum('STUDENT','INSTRUCTOR','ADMIN');default:'STUDENT'" json:"role"`
	StudentID         *string   `json:"studentId"` // Pointer allow null
	InstructorID      *string   `json:"instructorId"`
	Image             string    `gorm:"default:''" json:"image"`
	Points            int       `gorm:"default:0" json:"points"`
	TotalCourses      int       `gorm:"default:0" json:"totalCourses"`
	BadgesCount       int       `gorm:"column:badges;default:0" json:"badgesCount"` // Renamed field to avoid clash with Relation
	InstructorCourses int       `json:"instructorCourses"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`

	// Relations
	EnrolledCourses []UserCourse `gorm:"foreignKey:UserID" json:"enrolledCourses,omitempty"`
	UserBadges      []UserBadge  `gorm:"foreignKey:UserID" json:"userBadges,omitempty"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}
