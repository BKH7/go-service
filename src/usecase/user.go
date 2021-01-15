package usecase

import (
	"gitlab.com/bunlert274/go-service/src/domain"
	"gitlab.com/bunlert274/go-service/src/model"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase ...
func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{userRepo}
}

func (r *userUseCase) GetUser(user *model.User, id string) error {
	if err := r.userRepo.GetUser(user, id); err != nil {
		return err
	}
	return nil
}

func (r *userUseCase) GetAllUser() ([]model.User, error) {
	var users []model.User
	err := r.userRepo.GetAllUser(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userUseCase) CreateUser(user *model.User) error {
	if err := r.userRepo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (r *userUseCase) UpdateUser(user *model.User, id string) error {
	if err := r.userRepo.UpdateUser(user, id); err != nil {
		return err
	}
	return nil
}

func (r *userUseCase) DeleteUser(id string) error {
	if err := r.userRepo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
