package test

import (
	"geoproperty_be/config"
	"geoproperty_be/domain"
	area_respository "geoproperty_be/repository/area"
	area_usecase "geoproperty_be/usecase/area"

	property_respository "geoproperty_be/repository/property"
	property_usecase "geoproperty_be/usecase/property"

	query_usecase "geoproperty_be/usecase/query_agent"
	"testing"

	"gorm.io/gorm"
)

var (
	db                 *gorm.DB
	err                error
	propertyRepository domain.PropertyRepository
	propertyUseCase    domain.PropertyUsecase
)

func TestMain(m *testing.M) {
	if err := config.InitializeConfig("../../.env"); err != nil {
		panic(err)
	}

	// Connect to Database
	db, err = config.Connect()

	if err != nil {
		panic(err)
	}

	queryAgentConnection := config.GPTService()
	queryAgentInitialize := query_usecase.NewPromptServiceClient(queryAgentConnection)
	queryAgentUseCase := query_usecase.NewUseCase(queryAgentInitialize)

	// Initalize Resources
	areaRepository := area_respository.NewRepository(db)
	areaUseCase := area_usecase.NewUseCase(areaRepository)

	propertyRepository = property_respository.NewRepository(db)
	propertyUseCase = property_usecase.NewUseCase(propertyRepository, areaUseCase, *queryAgentUseCase)

	m.Run()

}
