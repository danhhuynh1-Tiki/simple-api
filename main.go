package main

import (
	_ "fmt"
	"api/app"
)
func main(){
	app.Run()
	// a := models.ConnectDB()
	// if a == nil{
	// 	fmt.Println("Failed")
	// }else{
	// 	fmt.Println("Successfull")
	// }
}