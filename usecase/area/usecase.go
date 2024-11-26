package area

import (
	"geoproperty_be/domain"
	"geoproperty_be/utils"

	"errors"

	"github.com/spatial-go/geoos/space"
	"gorm.io/gorm"
)

type UseCase struct {
	AreaRepository domain.AreaRepository
}

// GetAreaByGeom implements domain.AreaUsecase.
func (u *UseCase) GetAreaByGeom(geom space.Geometry) (*domain.Area, error) {
	geom_wkt, err := utils.DecodeGeomWKT(geom)

	if err != nil {
		return nil, err
	}

	area, err := u.AreaRepository.GetAreaByGeom(geom_wkt.(string))

	if err != nil {
		return nil, err
	}

	return area, nil
}

// Overlaps implements domain.AreaUsecase.
func (u *UseCase) Overlaps(geom space.Polygon) (bool, error) {
	geom_wkt, err := utils.DecodeGeomWKT(geom)
	if err != nil {
		return false, err
	}

	overlaps, err := u.AreaRepository.Overlaps(geom_wkt.(string))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return overlaps, nil
}

func NewUseCase(areaRepo domain.AreaRepository) domain.AreaUsecase {
	return &UseCase{
		AreaRepository: areaRepo,
	}
}
