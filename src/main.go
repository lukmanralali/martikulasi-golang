package main

import (
	"./controllers"
	"./controllers/v1"
	"./controllers/v2"
	"./database"
	"./middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rollbar/rollbar-go"
	"log"
	"os"
)

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        // os.Exit(1)
    }
	
	if os.Getenv("APP_ENV") == "production" {
		rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
		rollbar.SetEnvironment(os.Getenv("APP_ENV"))
		rollbar.WrapAndWait(startApp)
	} else {
		startApp()
	}

}

func startApp() {

	defaultMiddleware := middleware.DefaultMiddleware{}

	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())

	v1.V1UserControllerHandler(router)
	v1.V1AuthenticationControllerHandler(router)
	v2.V2UserControllerHandler(router)
	controllers.UrlShrotenerControllerHandler(router)

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)

	// run database migration
	database.Migrate()

	router.Run(serverString)

}
