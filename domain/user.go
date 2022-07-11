package domain

import (
	// "api/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type UserUsecase interface {
	FindUser(id string) (bool, *User)
	DeleteUser(id string) bool
	UpdateUser(user User, id string) bool
	AddUser(user User) bool
	GetUser() []User
}

type UserRepository interface {
	UpdateUser(user User, id string) bool
	DeleteUser(id string) bool
	FindUser(id string) *User
	AddUser(user User) bool
	GetUser() []User
}
