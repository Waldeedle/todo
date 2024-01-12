package main

import (
	"github.com/labstack/echo/v4"
	"github.com/waldeedle/todo/internal/database"
	"github.com/waldeedle/todo/internal/private_api"
	"github.com/waldeedle/todo/internal/public_api"
	"github.com/waldeedle/todo/internal/todos"
)

func main() {
	e := echo.New()
	e.Debug = true

	db, err := database.AddDB()
	if err != nil {
		panic(err)
	}

	todoRepository := todos.NewRepository(db)
	todoService := todos.NewService(todoRepository)
	publicAPI := public_api.New(nil, todoService)

	err = publicAPI.AddRoutes(e)
	if err != nil {
		panic(err)
	}

	err = private_api.AddRoutes(e)
	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":8100"))
}
