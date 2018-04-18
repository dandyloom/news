package models

import "github.com/globalsign/mgo/bson"

type Post struct {
	ID				bson.ObjectId	`bson:"_id" json:"id"`
	Title			string			`bson:"title" json:"title"`
	Description		string			`bson:"description" json:"description"`
	Link			string			`bson:"link" json:"link"`
}