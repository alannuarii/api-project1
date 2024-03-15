package main

import (
	"api-project1/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.POST("api/notulen", controllers.PostNotulen)

	r.GET("api/notulen", controllers.GetNotulen)
	r.GET("api/allnotulen", controllers.GetAllNotulen)

	r.PUT("api/fotonotulen", controllers.PostFotoNotulen)

	r.Run(":8888")
}