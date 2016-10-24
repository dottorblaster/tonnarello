package main

import "gopkg.in/mgo.v2/bson"

type Pasta struct {
	_id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Label string
	Content string
}
