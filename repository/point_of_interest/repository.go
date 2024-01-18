package point_of_interest

import (
	"geoproperty_be/domain"

	"github.com/spatial-go/geoos/space"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// GeoJSON implements domain.POIRepository.
func (r *Repository) GeoJSON(centerPoint space.Point) (*domain.GeoData, error) {
	var features []domain.Feature

	if err := r.DB.Model(&domain.PointOfInterest{}).Select("'Feature' as type, json_build_object('nama', nama, 'kategori', kategori, 'jarak',  ST_Distance ( geom, ref_geom ))::jsonb as properties, ST_SetSRID(geom,0)::jsonb as geometry").Joins("CROSS JOIN ( SELECT ST_MakePoint (?, ?) :: geography AS ref_geom ) AS r", centerPoint.X(), centerPoint.Y()).Where("ST_DWithin ( geom, ref_geom, ?, true)", 2000).Order("ST_Distance ( geom, ref_geom )").Find(&features).Error; err != nil {
		return nil, err
	}

	geoData := domain.GeoData{
		Type:     "FeatureCollection",
		Features: features,
	}

	return &geoData, nil
}

func NewRepository(db *gorm.DB) domain.POIRepository {
	return &Repository{
		DB: db,
	}
}
