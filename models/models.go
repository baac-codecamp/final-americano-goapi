package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type News struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	ImgURL    string             `json:"imgUrl" bson:"imgUrl,omitempty"`
	Desc      string             `json:"desc" bson:"desc,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
}



