package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"cookhub.com/app/api/v1/test"
	"cookhub.com/app/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var certFile = "/tmp/cert.pem"
var keyFile = "/tmp/key.pem"

func main() {
	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	_, err := db.InitStore()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	server.GET("/", func (context echo.Context) error {
		return context.HTML(http.StatusOK, fmt.Sprintf("Hello, CookHub!"))
	})

	server.GET("/v1/ping", test.HandleTest)

	server.GET("/v1/sum", test.HandleSum)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	//server.Logger.Fatal(server.Start(":" + httpPort))
	server.Logger.Fatal(server.StartTLS(":" + httpPort, certFile, keyFile))
}