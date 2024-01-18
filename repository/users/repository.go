package users

import (
	"geoproperty_be/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// Insert implements domain.UsersRepository.
func (r *Repository) Insert(user *domain.Users) (*domain.Users, error) {
	err := r.DB.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Find implements domain.UsersRepository.
func (r *Repository) Find(param map[string]any) (*[]domain.Users, error) {
	var users []domain.Users

	if err := r.DB.Where(param).Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func NewRepository(db *gorm.DB) domain.UsersRepository {
	return &Repository{
		DB: db,
	}
}
