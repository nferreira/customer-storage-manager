package model

import "time"

type Customer struct {
	Id             string         `json:"id" bson:"_id"`
	Name           string         `json:"name,omitempty" bson:"name,omitempty"`
}