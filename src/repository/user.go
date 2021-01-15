package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitlab.com/bunlert274/go-service/src/domain"
	"gitlab.com/bunlert274/go-service/src/model"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

// NewUserRepository ...
func NewUserRepository(conn *gorm.DB) domain.UserRepository {
	return &userRepository{conn}
}

func (r *userRepository) GetByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	if err := r.conn.WithContext(ctx).Table("users").Where(`userId = ?`, id).First(&user).Error; err != nil {
		logrus.Error(err)
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetDupplicate(ctx context.Context, user *model.User) bool {
	var u model.User
	r.conn.WithContext(ctx).Where(user.Username).Take(u)

	if &u != user {
		return true
	}
	return false
}

func (r *userRepository) IsExist(ctx context.Context, id string) bool {
	var user model.User
	r.conn.WithContext(ctx).Where(`userId = ?`, id).Take(user)
	if user.Username != "" {
		return true
	}
	return false
}

func (r *userRepository) Fetch(ctx context.Context) ([]model.User, error) {
	var user []model.User
	if err := r.conn.WithContext(ctx).Find(user).Error; err != nil {
		logrus.Error(err)
		return []model.User{}, err
	}
	return user, nil
}

func (r *userRepository) Store(ctx context.Context, user *model.User) error {
	if err := r.conn.WithContext(ctx).Create(user).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, id string, user *model.User) error {
	if err := r.conn.WithContext(ctx).Where("userId = ?", id).Updates(user).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	if err := r.conn.WithContext(ctx).Where("userId = ?", id).Updates(map[string]interface{}{
		"isDeleted": true,
	}).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
