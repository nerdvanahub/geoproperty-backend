package property

import (
	"context"
	"geoproperty_be/domain"

	query_agent "geoproperty_be/usecase/query_agent"

	"github.com/spatial-go/geoos/space"
)

type UseCase struct {
	PropertyRepository domain.PropertyRepository
	AreaUseCase        domain.AreaUsecase
	QueryAgentUseCase  query_agent.UseCase
}

// GetPropertyByPrompt implements domain.PropertyUsecase.
func (u *UseCase) GetPropertyByPrompt(query string) (*domain.GeoData, error) {
	// Get Query
	queryAgent, err := u.QueryAgentUseCase.GetQuery(context.TODO(), &query_agent.Prompt{
		Prompt: query,
	})

	if err != nil {
		return nil, err
	}

	// Generate Query
	generatedQuery, err := u.PropertyRepository.Generate(queryAgent.Response)

	if err != nil {
		return nil, err
	}

	// Find Property
	properties, err := u.PropertyRepository.Find(map[string]any{
		"id": generatedQuery,
	})

	if err != nil {
		return nil, err
	}

	var newProperties []domain.Property[space.Point, space.Polygon]
	for _, property := range *properties {
		// Maping Data
		result, err := property.MapGeom(property)

		if err != nil {
			return nil, err
		}

		newProperties = append(newProperties, *result)
	}

	var geoData domain.GeoData
	var propertiesEncoded []domain.Feature

	// Encode Properties
	for _, property := range newProperties {
		propertyEncoded, err := property.MapGeoJSON(property)

		if err != nil {
			return nil, err
		}

		propertiesEncoded = append(propertiesEncoded, *propertyEncoded)
	}

	geoData = domain.GeoData{
		Type:     "FeatureCollection",
		Features: propertiesEncoded,
	}

	return &geoData, nil
}

// GetByPoint implements domain.PropertyUsecase.
func (u *UseCase) GetByPoint(point space.Point) (*domain.GeoData, error) {
	area, err := u.AreaUseCase.GetAreaByGeom(point)

	if err != nil {
		return nil, err
	}

	properties, err := u.PropertyRepository.Find(map[string]any{
		"kelurahan": area.Kelurahan,
	})

	if err != nil {
		return nil, err
	}

	var newProperties []domain.Property[space.Point, space.Polygon]
	for _, property := range *properties {
		// Maping Data
		result, err := property.MapGeom(property)

		if err != nil {
			return nil, err
		}

		newProperties = append(newProperties, *result)
	}

	var geoData domain.GeoData
	var propertiesEncoded []domain.Feature

	// Encode Properties
	for _, property := range newProperties {
		propertyEncoded, err := property.MapGeoJSON(property)

		if err != nil {
			return nil, err
		}

		propertiesEncoded = append(propertiesEncoded, *propertyEncoded)
	}

	geoData = domain.GeoData{
		Type:     "FeatureCollection",
		Features: propertiesEncoded,
	}

	return &geoData, nil
}

// Delete implements domain.PropertyUsecase.
func (u *UseCase) Delete(uid string) error {
	if err := u.PropertyRepository.Delete(uid); err != nil {
		return err
	}

	return nil
}

// Insert implements domain.PropertyUsecase.
func (u *UseCase) Insert(property domain.Property[space.Point, space.Polygon]) (*domain.Property[space.Point, space.Polygon], error) {
	// Set UUID
	property.SetUID()

	// Parse WKT
	parsedWKTProperty, err := property.MapWKT(property)

	if err != nil {
		return nil, err
	}

	// Check Intersect
	area, err := u.AreaUseCase.GetAreaByGeom(property.Geometry)

	if err != nil {
		return nil, err
	}

	parsedWKTProperty.Kecamatan = area.Kecamatan
	parsedWKTProperty.Kelurahan = area.Kelurahan
	parsedWKTProperty.Kota = area.Kota

	// Insert Data
	newProperty, err := u.PropertyRepository.Insert(*parsedWKTProperty)

	if err != nil {
		return nil, err
	}

	// Parse Geometry
	result, err := (*newProperty).MapGeom(*newProperty)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindDetail implements domain.PropertyUsecase.
func (u *UseCase) FindDetail(uid string) (*domain.Property[space.Point, space.Polygon], error) {
	property, err := u.PropertyRepository.Find(map[string]any{
		"uuid": uid,
	})

	if err != nil {
		return nil, err
	}

	// Maping Data
	result, err := (*property)[0].MapGeom((*property)[0])

	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindAll implements domain.PropertyUsecase.
func (u *UseCase) FindAll(param map[string]any) (*[]domain.Property[space.Point, space.Polygon], error) {
	properties, err := u.PropertyRepository.Find(param)

	if err != nil {
		return nil, err
	}

	_, user_id_exist := param["user_id"]

	var newProperties []domain.Property[space.Point, space.Polygon]
	for index, property := range *properties {
		// Check if user_id exist and index > 100
		if user_id_exist && index > 100 {
			break
		}

		// Maping Data
		result, err := property.MapGeom(property)

		if err != nil {
			return nil, err
		}

		newProperties = append(newProperties, *result)
	}

	return &newProperties, nil
}

func NewUseCase(r domain.PropertyRepository, a domain.AreaUsecase, q query_agent.UseCase) domain.PropertyUsecase {
	return &UseCase{
		PropertyRepository: r,
		AreaUseCase:        a,
		QueryAgentUseCase:  q,
	}
}
