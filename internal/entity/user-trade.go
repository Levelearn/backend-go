package entity

import "time"

type UserTrade struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int       `json:"userId"`
	TradeID     int       `json:"tradeId"`
	PurchasedAt time.Time `gorm:"autoCreateTime" json:"purchasedAt"`
}
