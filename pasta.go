package main

import "gopkg.in/mgo.v2/bson"

// Pasta is an exported type that serves as a target for MongoDB
// document serialization/deserialization.
type Pasta struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Label   string
	Content string
}
