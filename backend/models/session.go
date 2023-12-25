package models

type Session struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username,omitempty" json:"username"`
}
