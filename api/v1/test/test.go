package test

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleTest(context echo.Context) error {
	return context.HTML(http.StatusOK, fmt.Sprintf("<h1>Test ping!</h1>"))
}

func HandleSum(context echo.Context) error {
	return context.HTML(http.StatusOK, fmt.Sprintf("<h1>42</h1>"))
}