package domain

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Users struct {
	ID        uint           `gorm:"primarykey"`
	Name      string         `gorm:"not null" validate:"required"`
	Email     string         `gorm:"unique;not null" validate:"required,email"`
	Password  string         `gorm:"not null" validate:"required"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (u *Users) TableName() string {
	return "users"
}

type UsersRepository interface {
	Find(param map[string]any) (*[]Users, error)
	Insert(user *Users) (*Users, error)
}

type UsersUseCase interface {
	Login(email string, password string) (*Token, error)
	Register(user *Users) (*Users, error)
	ExtractTokenGoogle(code string) (string, error)
}
