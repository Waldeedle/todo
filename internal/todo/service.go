package todo

import "database/sql"

type TodoService struct {
	db *sql.DB
}

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{
		db: db,
	}
}

func (t *TodoService) GetAllTodos() ([]Todo, error) {
	rows, err := t.db.Query("SELECT id, title, is_completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *TodoService) Create(title string) error {
	_, err := t.db.Exec("INSERT INTO todos (title, is_completed) VALUES ($1, $2)", title, false)
	if err != nil {
		return err
	}

	return nil
}
