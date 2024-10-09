package dto

type FollowRequest struct {
	FollowerID string `json:"follower_id"`
	FolloweeID string `json:"followee_id"`
}

type FilterFollow struct {
	ID         int64
	FollowerID string
	FolloweeID string
}
