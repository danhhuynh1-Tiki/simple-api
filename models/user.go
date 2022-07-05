package models

import(
	"time"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)
type User struct{
	
	ID string `json:"ID" bson:"ID"`
	Name string `json:"Name" bson:"Name"`
}
func ConnectDB() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
			return nil
	}else{
		return client
	}
}
func GetUser() []User{
	client := ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer client.Disconnect(ctx)

	example := client.Database("example")
	
	userc := example.Collection("user")

	cursor,err := userc.Find(ctx,bson.M{})

	defer cursor.Close(ctx)

	alluser := []User{}
	for cursor.Next(ctx){
		var user User
		cursor.Decode(&user)

		alluser = append(alluser,user)
	}
	if err != nil{
		return nil
	}else{
		return alluser
	}
}
func AddUser(id string,name string) bool{
	client := ConnectDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer client.Disconnect(ctx)

	example := client.Database("example")

	userc := example.Collection("user")

	res, err := userc.InsertOne(ctx,bson.D{
		{Key : "ID",Value:id},
		{Key:"Name",Value:name},
	})
	if err != nil{
		return false
	}else{
		fmt.Println(res)
		return true
	}
}