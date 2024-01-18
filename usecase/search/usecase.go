package search

import (
	"geoproperty_be/domain"
	"strings"
)

type UseCase struct {
	SearchRepository domain.SearchRepository
}

// GetAll implements domain.SearchUseCase.
func (r *UseCase) GetAll() (*[]domain.Search, error) {
	search, err := r.SearchRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return search, nil
}

// Search implements domain.SearchUseCase.
func (u *UseCase) Search(keyword string) (*[]domain.Search, error) {
	search, err := u.SearchRepository.Search(strings.ToLower(keyword))

	if err != nil {
		return nil, err
	}

	return search, nil
}

func NewUseCase(r domain.SearchRepository) domain.SearchUseCase {
	return &UseCase{
		SearchRepository: r,
	}
}
