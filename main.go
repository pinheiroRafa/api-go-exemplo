package main

import (
	"log"
	controller "main/src/controller"
	"main/src/utils"
	"os"

	db "main/src/repository"

	_ "main/docs" // gin-swagger middleware

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files" //
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           GO API
// @version         1.0
// @description     Fazendo uma API Restfull em GO
// @termsOfService  http://swagger.io/terms/

// @contact.name   Rafael
// @contact.url    https://www.linkedin.com/in/rafael-pg/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Param Content-Language   header string false "Idioma"
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.OpenDB(true)
	r := gin.Default()
	r.Use(utils.LangVerify)
	controller.UserController(r)
	controller.StatusController(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + os.Getenv("PORT"))
}
