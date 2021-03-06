// GENERATED BY goruda
// This file was generated automatically at
// 2020-07-17 07:53:27.959631 +0700 WIB m=+0.006716046

package articletest

import (
	"time"
)

type Article struct {
	Title     string                 `json:"title" bson:"title"`
	UpdatedAt time.Time              `json:"updated_at" bson:"updated_at"`
	Author    Author                 `json:"author" bson:"author"`
	CreatedAt time.Time              `json:"created_at" bson:"created_at"`
	Id        string                 `json:"id" bson:"id"`
	Publisher map[string]interface{} `json:"publisher" bson:"publisher"`
	Tag       Tag                    `json:"tag" bson:"tag"`
}
