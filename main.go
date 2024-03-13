package main

import (
	"api-project1/controllers"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.POST("api/notulen", controllers.PostNotulen)

	r.Run(":8888")
}