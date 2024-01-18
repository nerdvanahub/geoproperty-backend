package point_of_interest

import (
	"geoproperty_be/domain"

	"github.com/spatial-go/geoos/space"
)

type UseCase struct {
	POIRepository domain.POIRepository
}

// GeoJSON implements domain.POIUsecase.
func (u *UseCase) GeoJSON(centerPoint space.Point) (*domain.GeoData, error) {
	geoData, err := u.POIRepository.GeoJSON(centerPoint)

	if err != nil {
		return nil, err
	}

	return geoData, nil
}

func NewUseCase(poiRepo domain.POIRepository) domain.POIUsecase {
	return &UseCase{
		POIRepository: poiRepo,
	}
}
