package main

import (
    "github.com/gin-gonic/gin"
	"ClearFive/controller"
)

func main(){
    server:=gin.Default()
    server.POST("/getUserData",controller.GetUserDetails)
    server.Run(":8080")
}