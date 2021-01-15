package repository

import (
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

func (r *userRepository) GetUser(user *model.User, id string) error {
	if err := r.conn.Where("userId = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAllUser(user *[]model.User) error {
	if err := r.conn.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	if err := r.conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUser(user *model.User, id string) error {
	if err := r.conn.Where("userId = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(id string) error {
	if err := r.conn.Where("userId = ?", id).Updates(map[string]interface{}{
		"isDeleted": true,
	}).Error; err != nil {
		return err
	}
	return nil
}
