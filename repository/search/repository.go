package search

import (
	"geoproperty_be/domain"

	"gorm.io/gorm"
)

type Respository struct {
	DB *gorm.DB
}

// GetAll implements domain.SearchRepository.
func (r *Respository) GetAll() (*[]domain.Search, error) {
	var search []domain.Search

	err := r.DB.Raw(`SELECT CONCAT(name, ', ', kelurahan,', ', kecamatan,', ', kota) as name, ST_Centroid(geom) as center_point FROM streets LIMIT 100
	UNION
	SELECT CONCAT(kelurahan,', ', kecamatan,', ', kota) as name, ST_Centroid(geom) as center_point FROM administrative_area`).Scan(&search).Error

	if err != nil {
		return nil, err
	}

	// Encode WKB geom.
	for i := range search {
		err := search[i].EncodeGeom()

		if err != nil {
			return nil, err
		}
	}

	return &search, nil
}

// Search implements domain.SearchRepository.
func (r *Respository) Search(keyword string) (*[]domain.Search, error) {
	var search []domain.Search

	err := r.DB.Raw(`SELECT * FROM (SELECT CONCAT(name, ', ', kelurahan,', ', kecamatan,', ', kota) as name, ST_Centroid(geom) as center_point FROM streets WHERE LOWER(name) LIKE $1 OR LOWER(kelurahan) LIKE $1 OR LOWER(kecamatan) LIKE $1 OR LOWER(kota) LIKE $1
	UNION
	SELECT CONCAT(kelurahan,', ', kecamatan,', ', kota) as name, ST_Centroid(geom) as center_point FROM administrative_area WHERE LOWER(kelurahan) LIKE $1 OR LOWER(kecamatan) LIKE $1 OR LOWER(kota) LIKE $1) AS search LIMIT 100`, "%"+keyword+"%").Scan(&search).Error

	if err != nil {
		return nil, err
	}

	// Encode WKB geom.
	for i := range search {
		err := search[i].EncodeGeom()

		if err != nil {
			return nil, err
		}
	}

	return &search, nil
}

func NewRepository(db *gorm.DB) domain.SearchRepository {
	return &Respository{
		DB: db,
	}
}
