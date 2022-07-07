package repository

import (
	"api/domain"
	"api/repository"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
}

func GetUser() []domain.User {
	client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	defer client.Disconnect(ctx)

	example := client.Database("example")

	userc := example.Collection("user")

	cursor, err := userc.Find(ctx, bson.M{})

	defer cursor.Close(ctx)

	alluser := []domain.User{}
	for cursor.Next(ctx) {
		var user domain.User
		cursor.Decode(&user)

		alluser = append(alluser, user)
	}
	if err != nil {
		return nil
	} else {
		return alluser
	}
}
func AddUser(name string) bool {
	client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	defer client.Disconnect(ctx)

	example := client.Database("example")

	userc := example.Collection("user")

	res, err := userc.InsertOne(ctx, bson.D{
		{Key: "name", Value: name},
	})
	if err != nil {
		return false
	} else {
		fmt.Println(res)
		return true
	}
}

func UpdateUser(id string, name string) bool {
	client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	defer client.Disconnect(ctx)

	example := client.Database("example")

	user := example.Collection("user")

	new_id, _ := primitive.ObjectIDFromHex(id)

	res, err := user.UpdateOne(
		ctx,
		bson.M{"_id": new_id},
		bson.D{
			{"$set", bson.D{{"name", name}}},
		},
	)
	fmt.Println("udpate count : ", res.ModifiedCount)
	if err != nil {
		return false
	} else {
		return true
	}
}

func DeleteUser(id string) bool {
	client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	example := client.Database("example")

	user := example.Collection("user")

	defer client.Disconnect(ctx)

	fmt.Println("delete id : ", id)

	new_id, _ := primitive.ObjectIDFromHex(id)

	res, err := user.DeleteOne(ctx, bson.M{"_id": new_id})

	fmt.Println("delete count : ", res.DeletedCount)

	if err != nil || res.DeletedCount == 0 {
		return false
	} else {
		return true
	}
}

func FindUser(id string) *domain.User {
	client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	example := client.Database("example")

	user := example.Collection("user")

	defer client.Disconnect(ctx)

	new_id, _ := primitive.ObjectIDFromHex(id)

	var us domain.User

	err := user.FindOne(ctx, bson.M{"_id": new_id}).Decode(&us)
	if err != nil {
		return nil
	} else {
		return &us
	}
}
