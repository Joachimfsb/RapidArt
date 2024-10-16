package models

type Follow struct {
	FollowerUserId int `json:"follower_user_id"`
	FolloweeUserId int `json:"followee_user_id"`
}
