package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Caption     string        `json:"name" bson:"Caption"`
	Image_url   string        `json:"email" bson:"Image_url"`
	Timestamp   string        `json:"passwd" bson:"Timestamp"`
}

