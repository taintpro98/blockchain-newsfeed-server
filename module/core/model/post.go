package model

import "time"

type PostModel struct {
	ID        string     `json:"id,omitempty" gorm:"column:id;default:uuid_generate_v4()"`
	Title     string     `json:"title,omitempty" gorm:"column:title"`
	Content   string     `json:"content,omitempty" gorm:"column:content"`
	MediaURL  string     `json:"media_url" gorm:"column:media_url"`
	UserID    string     `json:"user_id" gorm:"column:user_id"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (PostModel) TableName() string {
	return "posts"
}
