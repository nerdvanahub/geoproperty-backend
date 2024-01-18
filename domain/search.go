package domain

import "geoproperty_be/utils"

type Search struct {
	Name        string `json:"name" gorm:"column:name"`
	CenterPoint any    `json:"center_point" gorm:"column:center_point"`
}

func (s *Search) EncodeGeom() error {
	s.CenterPoint = *s.CenterPoint.(*interface{})
	result, err := utils.EncodeWKBGeom(s.CenterPoint.(string))

	if err != nil {
		return err
	}

	s.CenterPoint = result

	return nil
}

type SearchRepository interface {
	Search(keyword string) (*[]Search, error)
	GetAll() (*[]Search, error)
}

type SearchUseCase interface {
	Search(keyword string) (*[]Search, error)
	GetAll() (*[]Search, error)
}
