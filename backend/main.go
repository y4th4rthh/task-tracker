package main

import(
	"github.com/gin-gonic/gin"
	"gotask-pro/routes"
	"gotask-pro/config"
)

func main(){
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}