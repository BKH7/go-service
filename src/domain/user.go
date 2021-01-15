package domain

import (
	"gitlab.com/bunlert274/go-service/src/model"
)

// UserUseCase ...
type UserUseCase interface {
	GetUser(t *model.User, id string) error
	GetAllUser() (t []model.User, err error)
	CreateUser(t *model.User) error
	UpdateUser(t *model.User, id string) error
	DeleteUser(id string) error
}

// UserRepository ...
type UserRepository interface {
	GetUser(t *model.User, id string) error
	GetAllUser(t *[]model.User) error
	CreateUser(t *model.User) error
	UpdateUser(t *model.User, id string) error
	DeleteUser(id string) error
}
