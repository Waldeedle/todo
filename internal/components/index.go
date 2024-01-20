package components

import (
	"github.com/a-h/templ"
	"github.com/waldeedle/todo/internal/models"
)

func Index(todos []*models.Todo) templ.Component {
	var titles []string
	for _, todo := range todos {
		titles = append(titles, *todo.Title)
	}
	return index(titles)
}
