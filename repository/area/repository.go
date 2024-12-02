package area

import (
	"geoproperty_be/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// GetAreaByGeom implements domain.AreaRepository.
func (r *Repository) GetAreaByGeom(geom string) (*domain.Area, error) {
	var area *domain.Area

	if err := r.DB.Model(&domain.Area{}).Where("ST_Intersects(geom, ST_GeomFromText(?, 4326))", geom).First(&area).Error; err != nil {
		return nil, err
	}

	return area, nil
}

// Overlaps implements domain.AreaRepository.
func (r *Repository) Overlaps(geom string) (bool, error) {
	var area *domain.Area
	var streets *domain.Streets

	if err := r.DB.Table("property").Where("ST_Intersects(geometry, ST_GeomFromText(?, 4326))", geom).First(&area).Error; err != nil {
		return false, err
	}

	if err := r.DB.Table("streets").Where("ST_Intersects(geom, ST_GeomFromText(?, 4326))", geom).First(&streets).Error; err != nil {
		return false, err
	}

	return area != nil || streets != nil, nil
}

func NewRepository(db *gorm.DB) domain.AreaRepository {
	return &Repository{
		DB: db,
	}
}
