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
	DB *mongo.Database
}

func NewMongoUserRepository(DB *mongo.Database) domain.UserRepository {
	return &mongoUserRepository{DB}
}

func (c *mongoUserRepository) GetUser() []domain.User {
	// DB := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// defer DB.Disconnect(ctx)

	// example := DB.Database("example")

	userc := c.DB.Collection("user")

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
func (c *mongoUserRepository) AddUser(name string) bool {
	// DB := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// defer DB.Disconnect(ctx)

	// example := DB.Database("example")

	userc := c.DB.Collection("user")

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

func (c *mongoUserRepository) UpdateUser(id string, name string) bool {
	// DB := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// defer DB.Disconnect(ctx)

	// example := DB.Database("example")

	user := c.DB.Collection("user")

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

func (c *mongoUserRepository) DeleteUser(id string) bool {
	// DB := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// example := DB.Database("example")

	user := c.DB.Collection("user")

	// defer DB.Disconnect(ctx)

	// fmt.Println("delete id context : ", c1.Value("id"))

	new_id, _ := primitive.ObjectIDFromHex(id)

	res, err := user.DeleteOne(ctx, bson.M{"_id": new_id})

	fmt.Println("delete count : ", res.DeletedCount)

	if err != nil || res.DeletedCount == 0 {
		return false
	} else {
		return true
	}
}

func (c *mongoUserRepository) FindUser(id string) *domain.User {
	// DB := repository.ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// example := DB.Database("example")

	user := c.DB.Collection("user")

	// defer DB.Disconnect(ctx)

	new_id, _ := primitive.ObjectIDFromHex(id)

	var us domain.User

	err := user.FindOne(ctx, bson.M{"_id": new_id}).Decode(&us)
	if err != nil {
		return nil
	} else {
		return &us
	}
}
