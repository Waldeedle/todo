package accounts

import (
	"database/sql"
	"fmt"

	"github.com/waldeedle/todo/internal/models"
)

type Repository interface {
	Create(email string) (*models.Account, error)
	GetById(id int64) (*models.Account, error)
	GetByEmail(email string) (*models.Account, error)
	Update(id int64, email string) (*models.Account, error)
	Delete(id int64) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const accountsTable = "accounts"

func (r *repository) Create(email string) (*models.Account, error) {
	result, err := r.db.Exec(fmt.Sprintf("INSERT INTO %s (email) VALUES (?)", accountsTable), email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Could not get last insert ID", err)
		return nil, err
	}

	return r.GetById(id)
}

func (r *repository) GetById(id int64) (*models.Account, error) {
	row := r.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = ?", accountsTable), id)
	return r.scanAccount(row)
}

func (r *repository) GetByEmail(email string) (*models.Account, error) {
	row := r.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE email = ?", accountsTable), email)
	return r.scanAccount(row)
}

func (r *repository) Update(id int64, email string) (*models.Account, error) {
	_, err := r.db.Exec(fmt.Sprintf("UPDATE %s SET email = ? WHERE id = ?", accountsTable), email, id)
	if err != nil {
		return nil, err
	}

	return r.GetById(id)
}

func (r *repository) Delete(id int64) error {
	_, err := r.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = ?", accountsTable), id)
	return err
}

func (r *repository) scanAccount(row *sql.Row) (*models.Account, error) {
	var account *models.Account
	err := row.Scan(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
