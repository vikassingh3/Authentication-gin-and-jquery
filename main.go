package main

import (
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/vikas/config"

	_ "github.com/vikas/docs"
	"github.com/vikas/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API
// @version 2.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
// @securityDefinitions.basic  BasicAuth
func main() {

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	app := gin.Default()
	app.Use(cors.Default())
	// app.Use(gin.Logger())

	godotenv.Load()
	app.Static("/data/", "./temp")

	url := ginSwagger.URL("http://localhost:10000/swagger/doc.json") // The url pointing to API definition
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	config.ConnectDB()

	//Routes
	routes.Routes(app)
	// Start server
	app.Run(":10000")
}
