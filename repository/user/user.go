package repository

import (
	"api/domain"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	client *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Collection) domain.UserRepository {
	return &mongoUserRepository{client}
}

func (c *mongoUserRepository) GetUser() []domain.User {
	// client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// defer client.Disconnect(ctx)

	// example := client.Database("example")

	// userc := example.Collection("user")

	cursor, err := c.client.Find(ctx, bson.M{})

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
func (c *mongoUserRepository) AddUser(name string) bool {
	// client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// defer client.Disconnect(ctx)

	// example := client.Database("example")

	// userc := example.Collection("user")

	res, err := c.client.InsertOne(ctx, bson.D{
		{Key: "name", Value: name},
	})
	if err != nil {
		return false
	} else {
		fmt.Println(res)
		return true
	}
}

func (c *mongoUserRepository) UpdateUser(id string, name string) bool {
	// client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// defer client.Disconnect(ctx)

	// example := client.Database("example")

	// user := example.Collection("user")

	new_id, _ := primitive.ObjectIDFromHex(id)

	res, err := c.client.UpdateOne(
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

func (c *mongoUserRepository) DeleteUser(id string) bool {
	// client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// example := client.Database("example")

	// user := example.Collection("user")

	// defer client.Disconnect(ctx)

	// fmt.Println("delete id context : ", c1.Value("id"))

	new_id, _ := primitive.ObjectIDFromHex(id)

	res, err := c.client.DeleteOne(ctx, bson.M{"_id": new_id})

	fmt.Println("delete count : ", res.DeletedCount)

	if err != nil || res.DeletedCount == 0 {
		return false
	} else {
		return true
	}
}

func (c *mongoUserRepository) FindUser(id string) *domain.User {
	// client := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// example := client.Database("example")

	// user := example.Collection("user")

	// defer client.Disconnect(ctx)

	new_id, _ := primitive.ObjectIDFromHex(id)

	var us domain.User

	err := c.client.FindOne(ctx, bson.M{"_id": new_id}).Decode(&us)
	if err != nil {
		return nil
	} else {
		return &us
	}
}
