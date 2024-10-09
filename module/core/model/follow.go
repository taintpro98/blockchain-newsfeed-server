package model

import (
	"time"
)

type FollowModel struct {
	ID         int64     `json:"id" gorm:"column:id"`
	FollowerID string    `json:"follower_id" gorm:"column:follower_id"`
	FolloweeID string    `json:"followee_id" gorm:"column:followee_id"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (FollowModel) TableName() string {
	return "follows"
}
