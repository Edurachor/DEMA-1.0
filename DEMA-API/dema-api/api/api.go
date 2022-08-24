package api

import (
	"apidema/api/controllers"
	"apidema/api/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {

	fmt.Println("Running...")
	listen(3000)
}

func listen(p int) {

	port := fmt.Sprintf(":%d", p)

	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.GET("/users", controllers.FindAllUsers)

	r.GET("/dashboard", middlewares.Auth())

	r.Run(port)
}
