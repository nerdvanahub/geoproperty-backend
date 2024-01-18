package domain

import "github.com/spatial-go/geoos/space"

type PointOfInterest struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"nama" gorm:"not null;column:nama"`
	Kategori string `json:"kategori" gorm:"not null;column:kategori"`
	Geom     any    `json:"geom" gorm:"not null;column:geom;type:geometry(Point,4326)"`
}

func (*PointOfInterest) TableName() string {
	return "point_of_interest"
}

type POIRepository interface {
	GeoJSON(centerPoint space.Point) (*GeoData, error)
}

type POIUsecase interface {
	GeoJSON(centerPoint space.Point) (*GeoData, error)
}
