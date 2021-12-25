package main

import (
	"fmt"
	"os"
	"log"
	"net/http"

	"cookhub.com/app/api/v1/test"
	"cookhub.com/app/db"
	"cookhub.com/app/api/v1/onboarding"
	"cookhub.com/app/third_party/gofirebase"
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
	firebaseAuth := auth.FirebaseAuthMiddleware { 
		Client: authClient,
	}

	_, err = db.InitStore()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	

	server.GET("/", func (context echo.Context) error {
		return context.HTML(http.StatusOK, fmt.Sprintf("Hello, CookHub!"))
	})

	server.GET("/v1/ping", firebaseAuth.HandleAuth(test.HandleTest))

	server.GET("/v1/sum", test.HandleSum)

	server.GET("/v1/onboarding", onboarding.GetOnboarding)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	//server.Logger.Fatal(server.Start(":" + httpPort))
	server.Logger.Fatal(server.StartTLS(":" + httpPort, certFile, keyFile))
}