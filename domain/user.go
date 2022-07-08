package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type UserUsecase interface {
	FindUser(id string) (bool, *User)
	DeleteUser(id string) bool
	UpdateUser(id string, name string) bool
	AddUser(name string) bool
	GetUser() []User
}

type UserRepository interface {
	UpdateUser(id string, name string) bool
	DeleteUser(id string) bool
	FindUser(id string) *User
	AddUser(name string) bool
	GetUser() []User
}
