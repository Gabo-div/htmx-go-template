package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"just-html/src/index"
	"log"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Static("/static", "static")

	index.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8081"))
}
