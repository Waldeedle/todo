//go:generate mockgen -destination=../mocks/todos/repository.go -package=mock_todos github.com/waldeedle/todo/internal/todos Repository

package todos

import (
	"database/sql"
	"time"

	"github.com/waldeedle/todo/internal/models"
)

type Repository interface {
	Create(title string) (*models.Todo, error)
	GetAll() ([]*models.Todo, error)
	GetById(id int) (*models.Todo, error)
	GetByAccountId(accountID int) (*models.Todo, error)
	//todo: potentially add partial title search
	GetByTitle(title string) (*models.Todo, error)
	GetByIsCompleted(isCompleted bool) ([]*models.Todo, error)
	Update(updatedTodo *models.Todo) (*models.Todo, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

const todosTable = "todos"

func (r *repository) Create(title string) (*models.Todo, error) {
	result, err := r.db.Exec("INSERT INTO todos (account_id, title, created_at, updated_at) VALUES (0, ?, ?, ?)", title, time.Now().Format(time.DateTime), time.Now().Format(time.DateTime))
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetById(int(id))
}

func (r *repository) GetAll() ([]*models.Todo, error) {
	rows, err := r.db.Query("SELECT * FROM todos ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}

	return r.scanTodos(rows)
}

func (r *repository) GetById(id int) (*models.Todo, error) {
	row := r.db.QueryRow("SELECT * FROM todos WHERE id = ?", id)
	return r.scanTodo(row)
}

func (r *repository) GetByAccountId(accountID int) (*models.Todo, error) {
	row := r.db.QueryRow("SELECT * FROM todos WHERE account_id = ?", accountID)
	return r.scanTodo(row)
}

func (r *repository) GetByTitle(title string) (*models.Todo, error) {
	row := r.db.QueryRow("SELECT * FROM todos WHERE title = ?", title)
	return r.scanTodo(row)
}

func (r *repository) GetByIsCompleted(isCompleted bool) ([]*models.Todo, error) {
	rows, err := r.db.Query("SELECT * FROM todos WHERE is_completed = ?", isCompleted)
	if err != nil {
		return nil, err
	}

	return r.scanTodos(rows)
}

func (r *repository) Update(updatedTodo *models.Todo) (*models.Todo, error) {
	_, err := r.db.Exec(`UPDATE todos
		SET
			account_id = ?,
			title = ?,
			is_completed = ?,
			is_deleted = ?,
			created_at = ?,
			updated_at = ?,
			deleted_at = ?
		WHERE id = ?`,
		updatedTodo.GetAccountID(),
		updatedTodo.GetTitle(),
		updatedTodo.GetIsCompleted(),
		updatedTodo.GetIsDeleted(),
		updatedTodo.GetCreatedAt(),
		updatedTodo.GetUpdatedAt(),
		updatedTodo.GetDeletedAt(),
		updatedTodo.GetID(),
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(*updatedTodo.GetID())
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func (r *repository) scanTodos(rows *sql.Rows) ([]*models.Todo, error) {
	var todos []*models.Todo

	for rows.Next() {
		todo := models.Todo{}
		err := rows.Scan(&todo.ID, &todo.AccountID, &todo.Title, &todo.IsCompleted, &todo.IsDeleted, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}

	return todos, nil
}

func (r *repository) scanTodo(row *sql.Row) (*models.Todo, error) {
	todo := models.Todo{}

	err := row.Scan(&todo.ID, &todo.AccountID, &todo.Title, &todo.IsCompleted, &todo.IsDeleted, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
