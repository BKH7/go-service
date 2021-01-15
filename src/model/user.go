package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User ...
type User struct {
	ID        string `gorm:"column:userId;primary_key;unique" json:"userId"`
	Username  string `gorm:"column:username;unique" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
	Email     string `gorm:"column:email;unique" json:"email"`
	Avatar    string `gorm:"column:avatar" json:"avatar"`
	IsDeleted bool   `gorm:"default:false;column:isDeleted" json:"isDeleted"`
	CreatedAt int64  `gorm:"column:createdAt;autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updatedAt;autoUpdateTime:milli" json:"updatedAt"`
}

// TableName use to specific table
func (m *User) TableName() string {
	return "users"
}

// BeforeCreate change id to uuid instead of autoincrement
func (m *User) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.New().String()
	m.Password = makePassword(m.Password)
	return nil
}

func makePassword(passwd string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(bytes)
}
