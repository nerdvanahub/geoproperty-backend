package domain

import (
	"encoding/json"
	"fmt"
	"geoproperty_be/utils"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/spatial-go/geoos/space"
)

type Property[C string | space.Point, G string | space.Polygon] struct {
	ID              int64           `json:"id" gorm:"primaryKey"`
	UUID            string          `json:"uuid" gorm:"unique"`
	UserID          int64           `json:"user_id" gorm:"not null"`
	User            Users           `json:"user" gorm:"foreignKey:UserID"`
	Images          []PropertyImage `json:"images" gorm:"foreignKey:PropertyID"`
	DeletedImage    []string        `json:"deleted_image" gorm:"-"`
	Address         string          `json:"address" gorm:"not null;type:text"`
	TitleAds        string          `json:"title_ads" gorm:"not null"`
	TypeAds         string          `json:"type_ads" gorm:"not null"`
	TypeProperty    string          `json:"type_property" gorm:"type:varchar(50);column:type_property"`
	Condition       string          `json:"condition" gorm:"not null"`
	Description     string          `json:"description" gorm:"not null;type:text"`
	Price           int64           `json:"price" gorm:"not null"`
	RentType        string          `json:"rent_type,omitempty" gorm:"type:varchar(50)"`
	BuildingType    string          `json:"building_type" gorm:"type:varchar(50)"`
	SurfaceArea     int64           `json:"surface_area" gorm:"null;"`
	BuildingArea    int64           `json:"building_area" gorm:"not null"`
	BathRooms       int64           `json:"bath_rooms" gorm:"not null"`
	BedRooms        int64           `json:"bed_rooms" gorm:"not null"`
	Floors          int64           `json:"floors" gorm:"not null"`
	ParkArea        int64           `json:"park_area" gorm:"not null"`
	Furniture       bool            `json:"furniture" gorm:"not null"`
	ElectricalPower int64           `json:"electrical_power" gorm:"not null"`
	Oriented        string          `json:"oriented,omitempty" gorm:"null;type:varchar(50)"`
	Certificate     string          `json:"certificate,omitempty" gorm:"type:varchar(30)"`
	FacilityInDoor  pq.StringArray  `json:"facility_in_door" gorm:"not null;type:text[]"`
	FacilityOutDoor pq.StringArray  `json:"facility_out_door" gorm:"not null;type:text[]"`
	FullName        string          `json:"full_name" gorm:"type:varchar(50)"`
	PhoneNumber     string          `json:"phone_number" gorm:"type:varchar(50)"`
	Email           string          `json:"email" gorm:"type:varchar(50)"`
	CenterPoint     C               `json:"center_point" gorm:"not null;type:geometry(Point,4326)"` // 4326 is SRID
	Geometry        G               `json:"geometry" gorm:"not null;type:geometry(Polygon,4326)"`   // 4326 is SRID
	Kelurahan       string          `json:"kelurahan" gorm:"column:kelurahan"`
	Kecamatan       string          `json:"kecamatan" gorm:"column:kecamatan"`
	Kota            string          `json:"kota" gorm:"column:kota"`
}

func (p *Property[C, G]) SetUID() {
	p.UUID = uuid.New().String()
}

func (p *Property[C, G]) MapGeom(property Property[string, string]) (*Property[space.Point, space.Polygon], error) {
	centerPointEncoded, err := utils.EncodeWKBGeom(property.CenterPoint)
	if err != nil {
		return nil, err
	}

	geometryEncoded, err := utils.EncodeWKBGeom(property.Geometry)

	if err != nil {
		return nil, err
	}

	newProperty := Property[space.Point, space.Polygon]{
		ID:              property.ID,
		UUID:            property.UUID,
		UserID:          property.UserID,
		User:            property.User,
		Images:          property.Images,
		Address:         property.Address,
		TitleAds:        property.TitleAds,
		TypeAds:         property.TypeAds,
		TypeProperty:    property.TypeProperty,
		Condition:       property.Condition,
		Description:     property.Description,
		Price:           property.Price,
		RentType:        property.RentType,
		BuildingType:    property.BuildingType,
		BuildingArea:    property.BuildingArea,
		SurfaceArea:     property.SurfaceArea,
		BathRooms:       property.BathRooms,
		BedRooms:        property.BedRooms,
		Floors:          property.Floors,
		ParkArea:        property.ParkArea,
		Furniture:       property.Furniture,
		ElectricalPower: property.ElectricalPower,
		Oriented:        property.Oriented,
		FacilityInDoor:  property.FacilityInDoor,
		FacilityOutDoor: property.FacilityOutDoor,
		FullName:        property.FullName,
		Email:           property.Email,
		PhoneNumber:     property.PhoneNumber,
		CenterPoint:     centerPointEncoded.(space.Point),
		Geometry:        geometryEncoded.(space.Polygon),
		Kelurahan:       property.Kelurahan,
		Kecamatan:       property.Kecamatan,
		Kota:            property.Kota,
	}

	return &newProperty, nil
}

func (p *Property[C, G]) MapWKT(property Property[space.Point, space.Polygon]) (*Property[string, string], error) {
	centerPointEncoded, err := utils.DecodeGeomWKT(property.CenterPoint)
	if err != nil {
		return nil, err
	}

	// Set SRD
	centerPointEncoded = fmt.Sprintf("SRID=4326;%s", centerPointEncoded)

	geometryEncoded, err := utils.DecodeGeomWKT(property.Geometry)
	if err != nil {
		return nil, err
	}

	// Set SRD
	geometryEncoded = fmt.Sprintf("SRID=4326;%s", geometryEncoded)

	newProperty := Property[string, string]{
		ID:              property.ID,
		UUID:            property.UUID,
		UserID:          property.UserID,
		User:            property.User,
		Images:          property.Images,
		Address:         property.Address,
		TitleAds:        property.TitleAds,
		TypeAds:         property.TypeAds,
		TypeProperty:    property.TypeProperty,
		Condition:       property.Condition,
		Description:     property.Description,
		Price:           property.Price,
		RentType:        property.RentType,
		BuildingType:    property.BuildingType,
		BuildingArea:    property.BuildingArea,
		SurfaceArea:     property.SurfaceArea,
		BathRooms:       property.BathRooms,
		BedRooms:        property.BedRooms,
		Floors:          property.Floors,
		ParkArea:        property.ParkArea,
		Furniture:       property.Furniture,
		ElectricalPower: property.ElectricalPower,
		Oriented:        property.Oriented,
		FacilityInDoor:  property.FacilityInDoor,
		FacilityOutDoor: property.FacilityOutDoor,
		FullName:        property.FullName,
		Email:           property.Email,
		PhoneNumber:     property.PhoneNumber,
		CenterPoint:     centerPointEncoded.(string),
		Geometry:        geometryEncoded.(string),
		Kelurahan:       property.Kelurahan,
		Kecamatan:       property.Kecamatan,
		Kota:            property.Kota,
	}

	return &newProperty, nil
}

func (*Property[C, G]) MapGeoJSON(property Property[space.Point, space.Polygon]) (*Feature, error) {
	var newProperty *Feature
	geometryEncoded, err := utils.EncodeGeomGeoJSON(property.Geometry)

	if err != nil {
		return nil, err
	}

	// Encode Properties to JSON
	propertyEncoded, err := json.Marshal(property)

	if err != nil {
		return nil, err
	}

	newProperty = &Feature{
		Type:       "Feature",
		Properties: propertyEncoded,
		Geometry:   geometryEncoded,
	}

	return newProperty, nil
}

func (*Property[C, G]) TableName() string {
	return "property"
}

type PropertyImage struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	PropertyID int64  `json:"property_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Image      string `json:"image" gorm:"not null"`
}

func (*PropertyImage) TableName() string {
	return "property_image"
}

type PropertyRepository interface {
	Insert(property Property[string, string]) (*Property[string, string], error)
	Find(param map[string]any) (*[]Property[string, string], error)
	Update(property Property[string, string]) (*Property[string, string], error)
	Generate(query string) ([]int, error)
	Delete(id int) error
	FinByPolygon(polygon string) (*[]Property[string, string], error)
}

type PropertyUsecase interface {
	FindAll(param map[string]any) (*[]Property[space.Point, space.Polygon], error)
	FindDetail(uid string) (*Property[space.Point, space.Polygon], error)
	Insert(property Property[space.Point, space.Polygon]) (*Property[space.Point, space.Polygon], error)
	Delete(uid string) error
	GetByGeom(types string, point space.Point, polygon space.Polygon) (*GeoData, error)
	GetPropertyByPrompt(query string) (*GeoData, error)
	Update(property Property[space.Point, space.Polygon]) (*Property[space.Point, space.Polygon], error)
}
