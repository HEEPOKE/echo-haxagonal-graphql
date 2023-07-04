package models

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Username  string    `bson:"username" json:"username"`
	Email     string    `bson:"email" json:"email" index:"unique"`
	Password  *string   `bson:"password,omitempty" json:"password"`
	Tel       string    `bson:"tel" json:"tel"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
