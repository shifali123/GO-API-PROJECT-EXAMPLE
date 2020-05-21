package models

import (
	"github.com/globalsign/mgo/bson"

)

type FastUser struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
FirstName  string        `json:"first_name" bson:"first_name"`
MiddleName  string        `json:"middle_name" bson:"middle_name"`
LastName  string        `json:"last_name" bson:"last_name"`
Address  AddressInfo         `json:"address" bson:"address"`
Email string 						`json:"email" bson:"email"`
Phone string						`json:"phone" bson:"phone"`
	//Address     AddressInfo   `json: "address" bson: "address"`
	StripeID    string        `json:"stripe_id" bson:"stripe_id"`
	BrainTreeID string        `json:"brain_tree_id" bson:"brain_tree_id"`
	CreatedAt int64       `json:"created_at" bson:"created_at"`
}

type AddressInfo struct {
	Line1  string        `json:"line_1" bson:"line_1"`
	Line2  string        `json:"line_2" bson:"line_2"`
	City  string        `json:"city" bson:"city"`
	SubDivsion  string        `json:"subdivision" bson:"subdivision"`
	PostalCode  string        `json:"postal_code" bson:"postal_code"`
	Country string  `json:"country" bson:"country"`
}

// CollectionName collection name for MongoDB.
func (fastuser *FastUser) CollectionName() string {
	return "fastuseraccount"
}

// ObjectID is the unique ID for the Fast User

//func (fastuser *FastUser) ObjectID() interface{} {
//	return fastuser.ID
//}
