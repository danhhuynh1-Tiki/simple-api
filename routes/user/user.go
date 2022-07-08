package routes

import (
	"api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(route *gin.Engine, u domain.UserUsecase) {
	userHandler := &UserHandler{userUsecase: u}

	// Simple Api : v1
	v1 := route.Group("/v1")
	{
		v1.GET("/user", userHandler.AllUser)
		v1.POST("/add_user", userHandler.AddUser)
		v1.PUT("/update_user", userHandler.UpdateUser)
		v1.DELETE("/delete_user", userHandler.DeleteUser)
		v1.GET("/find_user", userHandler.FindUser)
	}
}

func (u *UserHandler) AllUser(c *gin.Context) {
	// client := db.ConnectDB()

	alluser := u.userUsecase.GetUser()
	if alluser == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Fail",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data ":  alluser,
	})
}

func (u *UserHandler) AddUser(c *gin.Context) {

	name := c.Query("name")

	check := u.userUsecase.AddUser(name)

	if check == true {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Successfull",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Fail",
		})
	}

}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	check := u.userUsecase.UpdateUser(id, name)

	if check == true {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Successful",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Fail",
		})
	}

}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Query("id")

	check := u.userUsecase.DeleteUser(id)

	if check == true {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Successful",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Fail",
		})
	}
}

func (u *UserHandler) FindUser(c *gin.Context) {
	id := c.Query("id")

	check, us := u.userUsecase.FindUser(id)

	if check == false && us == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,
			"data": us,
		})
	}
}
