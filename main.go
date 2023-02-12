package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"database/sql"

	"cookhub.com/app/api/v1/test"
	"cookhub.com/app/db"
	"cookhub.com/app/api/v1/onboarding"
	"cookhub.com/app/api/v1/recipes"
	"cookhub.com/app/third_party/gofirebase"
	"cookhub.com/app/models"
	auth "cookhub.com/app/middleware/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var certFile = "/tmp/server.crt"
var keyFile = "/tmp/server.key"

func main() {
	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.Static("/static", "assets")

	authClient, err := gofirebase.SetupAuth()
	if err != nil {
		log.Fatalf("failed to initialize Firebase: %s", err)
	}
	authMiddleware := auth.InitAuthMiddleware(authClient)

	var database *sql.DB
	database, err = db.InitStore()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	server.GET("/", func (context echo.Context) error {
		return context.HTML(http.StatusOK, fmt.Sprintf("Hello, CookHub!"))
	})

	server.GET("/v1/ping", authMiddleware.HandleAuth(test.HandleTest))

	server.GET("/v1/sum", test.HandleSum)

	server.GET("/v1/onboarding", func (context echo.Context) error {
		return onboarding.GetOnboarding(context, models.InitOnboarding(database))
	})

	server.GET("/v1/recipes", authMiddleware.HandleAuth(func (context echo.Context) error {
		return recipes.GetUserFeedRecipes(context, models.InitRecipes(database), models.InitUsers(database))	
	}))

	server.GET("/v1/recipe", authMiddleware.HandleAuth(func (context echo.Context) error {
		return recipes.GetRecipe(context, models.InitRecipes(database), models.InitUsers(database))
	}))

	server.GET("/userAgreement", func (context echo.Context) error {
		return context.HTML(http.StatusOK, "<h1>User agreement</h1><p>Will be there one day...</p>")
	})
	server.GET("/privacy", func (context echo.Context) error {
		return context.HTML(http.StatusOK, "<h1>Privacy</h1><p>Will be there one day...</p>")
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	//server.Logger.Fatal(server.Start(":" + httpPort))
	server.Logger.Fatal(server.StartTLS(":" + httpPort, certFile, keyFile))
}