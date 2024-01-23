// package public_api

// import (
// 	"github.com/labstack/echo/v4"
// 	"github.com/waldeedle/todo/internal/todos"
// 	"github.com/waldeedle/todo/internal/accounts"
// )

// type TodoHandler struct {
// 	todos todos.Repository
// 	accounts accounts.Repository
// }

// func NewTodoHandler(e *echo.Echo, todoService todos.Repository, accounts accounts.Repository) {
// 	todoHandler := &TodoHandler{
// 		todos:    todoService,
// 		accounts: accounts,
// 	}

// 	e.POST("/create", todoHandler.CreateHandler)
// 	e.GET("/todo", todoHandler.GetAllHandler)
// 	e.GET("/todo/:id", todoHandler.GetHandler)
// 	e.POST("/todo/:id", todoHandler.UpdateHandler)
// 	e.DELETE("/todo/:id", todoHandler.DeleteHandler)
// }

// func HTML(c echo.Context, comp templ.Component) error {
// 	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
// 	return comp.Render(c.Request().Context(), c.Response().Writer)
// }

// type CreateTodoRequest struct {
// 	Title string `json:"title"`
// }

// type GetTodoRequest struct {
// 	ID int `json:"id"`
// }

// type UpdateTodoRequest struct {
// 	ID        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

// type DeleteTodoRequest struct {
// 	ID int `json:"id"`
// }

// func (t *TodoHandler) CreateHandler(c echo.Context) error {
// 	params := CreateTodoRequest{}
// 	if err := c.Bind(params); err != nil {
// 		return c.String(http.StatusBadRequest, "bad request")
// 	}
// 	_, err := t.todoService.Create(params.Title)
// 	if err != nil {
// 		return err
// 	}

// 	return HTML(c, templates.Success())
// }

// func (t *TodoHandler) GetAllHandler(c echo.Context) error {
// 	_, err := t.todoService.GetAll()
// 	if err != nil {
// 		return err
// 	}

// 	return HTML(c, templates.Page())
// }

// func (t *TodoHandler) GetHandler(c echo.Context) error {
// 	params := GetTodoRequest{}
// 	if err := c.Bind(params); err != nil {
// 		return c.String(http.StatusBadRequest, "bad request")
// 	}
// 	_, err := t.todoService.GetById(params.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return HTML(c, templates.Page())
// }

// func (t *TodoHandler) UpdateHandler(c echo.Context) error {
// 	params := UpdateTodoRequest{}
// 	if err := c.Bind(params); err != nil {
// 		return c.String(http.StatusBadRequest, "bad request")
// 	}
// 	_, err := t.todoService.Update(params.ID, params.Title, params.Completed)
// 	if err != nil {
// 		return err
// 	}

// 	return HTML(c, templates.Page())
// }

// func (t *TodoHandler) DeleteHandler(c echo.Context) error {
// 	params := DeleteTodoRequest{}
// 	if err := c.Bind(params); err != nil {
// 		return c.String(http.StatusBadRequest, "bad request")
// 	}
// 	err := t.todoService.Delete(params.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return HTML(c, templates.Page())
// }
