package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Created  time.Time          `bson:"created" json:"created"`
	Updated  time.Time          `bson:"updated" json:"updated"`
}



func (u *User ) ToJson() *User {
	// convert user data structure 
	// to json data struct 

	return nil 
}




type GetUserRequest struct {
	Id string 	`json:"_id"`
}

type ListUserRequest struct {

}


type DeleteUsersRequest struct {
	Id string `json:"_id"`
}

type DeleteUsersResponse struct {
	Id string `json:"_id"`
}
