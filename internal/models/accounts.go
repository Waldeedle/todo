package models

import "time"

type Account struct {
	ID        *int
	FirstName *string
	LastName  *string
	Email     *string
	IsDeleted *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (a *Account) GetID() *int {
	if a == nil {
		return nil
	}

	return a.ID
}

func (a *Account) GetFirstName() *string {
	if a == nil {
		return nil
	}

	return a.FirstName
}

func (a *Account) GetLastName() *string {
	if a == nil {
		return nil
	}

	return a.LastName
}

func (a *Account) GetEmail() *string {
	if a == nil {
		return nil
	}

	return a.Email
}

func (a *Account) GetIsDeleted() *bool {
	if a == nil {
		return nil
	}

	return a.IsDeleted
}
