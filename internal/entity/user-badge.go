package entity

import "time"

type UserBadge struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int       `json:"userId"`
	BadgeID     int       `json:"badgeId"`
	IsPurchased bool      `gorm:"default:false" json:"isPurchased"`
	AwardedAt   time.Time `gorm:"autoCreateTime" json:"awardedAt"`

	Badge Badge `gorm:"foreignKey:BadgeID" json:"badge"`
}
