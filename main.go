package main

import (
	"fmt"
	"os"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var certFile = "/tmp/cert.pem"
var keyFile = "/tmp/key.pem"

func main() {
	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.GET("/", func (context echo.Context) error {
		return context.HTML(http.StatusOK, fmt.Sprintf("Hello, CookHub!"))
	})

	server.GET("/v1/ping", func (context echo.Context) error {
		return context.HTML(http.StatusOK, fmt.Sprintf("<h1>Test ping!</h1>"))
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	//server.Logger.Fatal(server.Start(":" + httpPort))
	server.Logger.Fatal(server.StartTLS(":" + httpPort, certFile, keyFile))
}