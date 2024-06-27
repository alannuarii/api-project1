package main

import (
	"api-project1/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.Static("/photo", "./static/img")

	r.POST("api/notulen", controllers.PostNotulen)

	r.GET("api/notulen", controllers.GetNotulen)
	r.GET("api/notulen/:kode", controllers.GetNotulen)
	r.GET("api/allnotulen", controllers.GetAllNotulen)
	r.GET("api/alldaftarhadir", controllers.GetAllAbsensi)
	r.GET("api/daftarhadir/:kode", controllers.GetAbsensi)

	r.PUT("api/fotonotulen", controllers.PostFotoNotulen)

	r.DELETE("api/notulen", controllers.DeleteNotulen)

	r.Run(":8888")
}