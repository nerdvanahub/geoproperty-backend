package domain

import "github.com/spatial-go/geoos/space"

type Area struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Kelurahan string `json:"kelurahan" gorm:"column:kelurahan"`
	Kecamatan string `json:"kecamatan" gorm:"column:kecamatan"`
	Kota      string `json:"kota" gorm:"column:kota"`
	Geom      any    `json:"geom" gorm:"type:geometry(MultiPolygon,4326)"`
}

func (*Area) TableName() string {
	return "administrative_area"
}

type AreaRepository interface {
	GetAreaByGeom(geom string) (*Area, error)
	Overlaps(geom string) (bool, error)
}

type AreaUsecase interface {
	GetAreaByGeom(geom space.Geometry) (*Area, error)
	Overlaps(geom space.Polygon) (bool, error)
}
