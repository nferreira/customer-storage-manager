package model

type Customer struct {
	Id             string         `json:"id" bson:"_id"`
	Name           string         `json:"name,omitempty" bson:"name,omitempty"`
	Ssn           string          `json:"ssn,omitempty" bson:"ssn,omitempty"`
}
