package domain

import (
	"context"

	"gitlab.com/bunlert274/go-service/src/model"
)

// UserUseCase ...
type UserUseCase interface {
	GetByID(ctx context.Context, t *model.User, id string) error
	// Fetch() (t []model.User, err error)
	// Store(t *model.User) error
	// Update(t *model.User, id string) error
	// Delete(id string) error
}

// UserRepository ...
type UserRepository interface {
	GetByID(ctx context.Context, id string, user *model.User) error
	GetDupplicate(ctx context.Context, user *model.User) bool
	IsExist(ctx context.Context, id string) bool
	Fetch(ctx context.Context, user *[]model.User) error
	Store(ctx context.Context, user *model.User) error
	Update(ctx context.Context, id string, user *model.User) error
	Delete(ctx context.Context, id string) error
}
