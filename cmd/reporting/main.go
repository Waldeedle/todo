package main

import (
	"github.com/labstack/echo/v4"
	db "github.com/waldeedle/reporting/internal/database/sqlite"
	"github.com/waldeedle/reporting/internal/handlers"
	"github.com/waldeedle/reporting/internal/private_api"
	"github.com/waldeedle/reporting/internal/public_api"
	"github.com/waldeedle/reporting/internal/todo"
)

func main() {
	e := echo.New()
	e.Debug = true

	db, err := db.AddDB()
	if err != nil {
		panic(err)
	}

	todoService := todo.NewTodoService(db)
	//need to refactor to move this
	handlers.NewTodoHandler(e, todoService)

	err = public_api.AddRoutes(e)
	if err != nil {
		panic(err)
	}

	err = private_api.AddRoutes(e)
	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":3000"))
}
