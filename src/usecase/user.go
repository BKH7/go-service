package usecase

import (
	"context"
	"time"

	"gitlab.com/bunlert274/go-service/src/domain"
	"gitlab.com/bunlert274/go-service/src/model"
)

type userUseCase struct {
	userRepo domain.UserRepository
	timeout  time.Duration
}

// NewUserUseCase ...
func NewUserUseCase(userRepo domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{userRepo, timeout}
}

func (r *userUseCase) GetByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	ctx, cancel := context.WithTimeout(ctx, r.timeout)

	defer cancel()

	user, err := r.userRepo.GetByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// func (r *userUseCase) GetAllUser() ([]model.User, error) {
// 	var users []model.User
// 	err := r.userRepo.GetAllUser(&users)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

func (r *userUseCase) Store(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	if err := r.userRepo.Store(ctx, user); err != nil {
		return err
	}
	return nil
}

// func (r *userUseCase) UpdateUser(user *model.User, id string) error {
// 	if err := r.userRepo.UpdateUser(user, id); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *userUseCase) DeleteUser(id string) error {
// 	if err := r.userRepo.DeleteUser(id); err != nil {
// 		return err
// 	}
// 	return nil
// }
