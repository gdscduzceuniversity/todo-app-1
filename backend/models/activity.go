package models

import "time"

type Activity struct {
	ID     string    `bson:"_id,omitempty" json:"id"`
	Time   time.Time `bson:"time,omitempty" json:"time"`
	TaskID string    `bson:"taskID,omitempty" json:"taskID"`
	Type   int       `bson:"type,omitempty" json:"type"` // 0: status change, 1: title, description change, 2: due date change
	Old    string    `bson:"old,omitempty" json:"old"`
	New    string    `bson:"new,omitempty" json:"new"`
}
