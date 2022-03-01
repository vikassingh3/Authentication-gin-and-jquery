package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikas/auth"
	"github.com/vikas/config"
	"github.com/vikas/controllers"
	"github.com/vikas/middleware"
)

// func for describe group of private routes.

func Routes(router *gin.Engine) {

	redisClient := config.NewRedisDB()

	var rd = auth.NewAuth(redisClient)
	var tk = auth.NewToken()
	var service = controllers.NewProfile(rd, tk)

	/* router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	})) */

	{

		// router for POST method
		router.POST("api/login", gin.BasicAuth(gin.Accounts{"admin": "admin"}), service.Login)
		// router for  POST method
		router.POST("api/forgotPass", gin.BasicAuth(gin.Accounts{"admin": "admin"}), service.ForgotPass)
		// router for POST method
		router.POST("api/register", gin.BasicAuth(gin.Accounts{"admin": "admin"}), controllers.RegisterAdmin)

		{
			// Create routes group.
			app1 := router.Group("api/admin")
			// router for POST method
			app1.POST("/userlogin", gin.BasicAuth(gin.Accounts{"admin": "admin"}), service.UserLogin)
			app1.POST("/resetPass", gin.BasicAuth(gin.Accounts{"admin": "admin"}), service.ResetPass)

			// router for POST method
			app1.POST("/logout", gin.BasicAuth(gin.Accounts{"admin": "admin"}), middleware.TokenAuthMiddleware(), service.Logout)

			// router for GET method
			app1.GET("/allUser", gin.BasicAuth(gin.Accounts{"admin": "admin"}), middleware.TokenAuthMiddleware(), controllers.GetAllUsers)
			// router for POST method
			app1.POST("/createUser", controllers.CreateUser)
			// router for GET method
			app1.GET("/get/:id", gin.BasicAuth(gin.Accounts{"admin": "admin"}), middleware.TokenAuthMiddleware(), controllers.GetUser)
			// router for PUT method
			app1.PUT("/update/:id", gin.BasicAuth(gin.Accounts{"admin": "admin"}), middleware.TokenAuthMiddleware(), controllers.UpdateUser)
			// router for DELETE method
			app1.DELETE("/delete/:id", gin.BasicAuth(gin.Accounts{"admin": "admin"}), middleware.TokenAuthMiddleware(), service.DeleteUser)

		}
	}

}

// export PATH = $PATH:&HOME/go/bin
