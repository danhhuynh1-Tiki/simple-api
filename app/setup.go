package app

import (
	"api/repository"

	userRepo "api/repository/user"
	userHandl "api/routes/user"
	userUsercase "api/usecase"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	app := SetupRoute()
	app.Run()
}

func SetupRoute() *gin.Engine {
	app := gin.Default()
	client := repository.ConnectDB()

	userRepository := userRepo.NewMongoUserRepository(client)
	userUsecase := userUsercase.NewUserUsecase(userRepository)
	userHandl.NewUserHandler(app, userUsecase)

	return app
}
