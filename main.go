package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
)

func main() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("byandev %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.Use(cors.AllowAll())

	public := router.Group("/api/v1")
	{
		public.POST("/login_admin", controllers.LoginAdmin)
		public.POST("/login_petugas", controllers.LoginPetugas)
		public.POST("/login_users", controllers.LoginUsers)

		//public.POST("/create_akun", controllers.CreateUsers)
		//
		//public.GET("/akun_list", controllers.GetUsers)
		//public.GET("/akun_detail", controllers.GetUserDetail)
		//public.PUT("/akun_update", controllers.UpdateUser)
		//
		//public.GET("/rekom_list", controllers.GetRekom)
		//public.GET("/rekom_detail", controllers.GetRekomDetail)
		//public.POST("/rekom_create", controllers.CreateRekomendasi)
		//public.PUT("/rekom_update", controllers.UpdateRekom)
		//
		//public.GET("/privileges_role", controllers.GetRole)
		//public.GET("/privileges_type", controllers.GetType)

		//public.POST("/create_akun", controllers.CreateUsers)
		//public.GET("/akun/list", controllers.GetUsers)
		//public.GET("/akun/detail", controllers.GetUserDetail)
		//public.PUT("/akun/update", controllers.UpdateUser)
	}
	v1 := router.Group("/api/v1")
	//v1.Use(authentications.Auth)
	{
		v1.POST("/create_akun", controllers.CreateUsers)
		v1.GET("/search-users", controllers.SearchUser)
		v1.GET("/foto", controllers.GetFoto)
		v1.GET("/akun", controllers.GetUsers)
		v1.GET("/akun_list_jamaah", controllers.GetJamaah)
		v1.GET("/akun_list_petugas", controllers.GetPetugas)
		v1.GET("/akun_detail", controllers.GetUserDetail)
		v1.PUT("/akun_update", controllers.UpdateUser)
		v1.GET("/foto-rekomendasi", controllers.GetRekomFoto)
		v1.GET("/rekom_list", controllers.GetRekom)
		v1.GET("/rekom_detail", controllers.GetRekomDetail)
		v1.POST("/rekom_create", controllers.CreateRekomendasi)
		v1.PUT("/rekom_update", controllers.UpdateRekom)
		v1.GET("/privileges_role", controllers.GetRole)
		v1.GET("/privileges_type", controllers.GetType)
	}
	router.Run(":5000")
}
