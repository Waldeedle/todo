package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/waldeedle/todo/internal/templates"
	"github.com/waldeedle/todo/internal/todo"
)

type TodoHandler struct {
	todoService *todo.TodoService
}

func NewTodoHandler(e *echo.Echo, todoService *todo.TodoService) {
	todoHandler := &TodoHandler{
		todoService: todoService,
	}

	e.GET("/todo", todoHandler.GetTodoHandler)
	e.POST("/create", todoHandler.PostTodoHandler)
}

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func (t *TodoHandler) GetTodoHandler(c echo.Context) error {
	_, err := t.todoService.GetAllTodos()
	if err != nil {
		return err
	}

	return HTML(c, templates.Page())
}

func (t *TodoHandler) PostTodoHandler(c echo.Context) error {
	title := c.FormValue("title")
	err := t.todoService.Create(title)
	if err != nil {
		return err
	}

	return HTML(c, templates.Success())
}
