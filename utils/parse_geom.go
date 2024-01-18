package utils

import (
	"bytes"

	"github.com/spatial-go/geoos/geoencoding"
	"github.com/spatial-go/geoos/space"
)

func EncodeWKBGeom(wkb string) (space.Geometry, error) {
	buf := new(bytes.Buffer)
	buf.Write([]byte(wkb))
	geom, err := geoencoding.Read(buf, geoencoding.WKB)

	if err != nil {
		return nil, err
	}

	return geom.Geom(), nil
}

func DecodeGeomWKT(geom space.Geometry) (any, error) {
	buf := new(bytes.Buffer)
	err := geoencoding.Write(buf, geom, geoencoding.WKT)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func EncodeGeomGeoJSON(geom space.Geometry) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := geoencoding.Write(buf, geom, geoencoding.GeoJSON)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
