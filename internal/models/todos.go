package models

import "time"

type Todo struct {
	ID          *int64
	AccountID   *int64
	Title       *string
	IsCompleted *bool
	//todo: potentially add an archive feature?? also perma delete option
	IsDeleted *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (t *Todo) GetID() *int64 {
	if t == nil {
		return nil
	}

	return t.ID
}

func (t *Todo) GetAccountID() *int64 {
	if t == nil {
		return nil
	}

	return t.AccountID
}

func (t *Todo) GetTitle() *string {
	if t == nil {
		return nil
	}

	return t.Title
}

func (t *Todo) GetIsCompleted() *bool {
	if t == nil {
		return nil
	}

	return t.IsCompleted
}

func (t *Todo) GetIsDeleted() *bool {
	if t == nil {
		return nil
	}

	return t.IsDeleted
}
