package models

type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username,omitempty" json:"username"`
	Password string `bson:"password,omitempty" json:"password"`
}
