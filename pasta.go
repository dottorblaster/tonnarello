package main

import "gopkg.in/mgo.v2/bson"

type Pasta struct {
	_id bson.ObjectId
	Label string
	Content string
}
