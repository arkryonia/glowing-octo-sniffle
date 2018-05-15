package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionArticle prend le nom de collection `articles`
	CollectionArticle = "articles"
)

// Article est un model
type Article struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title" form:"title" binding:"required" bson:"title"`
	Body      string        `json:"body" form:"body" binding:"required" bson:"body"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}
