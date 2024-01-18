package command

import (
	"context"
	"geoproperty_be/config"
	"geoproperty_be/delivery/http/controllers"
	"geoproperty_be/delivery/http/routes"

	area_respository "geoproperty_be/repository/area"
	area_usecase "geoproperty_be/usecase/area"

	assets_usecase "geoproperty_be/usecase/assets"

	poi_respository "geoproperty_be/repository/point_of_interest"
	poi_usecase "geoproperty_be/usecase/point_of_interest"

	property_respository "geoproperty_be/repository/property"
	property_usecase "geoproperty_be/usecase/property"

	search_respository "geoproperty_be/repository/search"
	search_usecase "geoproperty_be/usecase/search"

	users_respository "geoproperty_be/repository/users"
	users_usecase "geoproperty_be/usecase/users"

	query_agent_usecase "geoproperty_be/usecase/query_agent"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
)

var (
	ip, port string

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run Server",
		Long:  `Run Server with IP and Port`,
		Run: func(cmd *cobra.Command, args []string) {

			// Connect to Database
			db, err := config.Connect()

			if err != nil {
				panic(err)
			}

			// Connect to MinIO
			minioClient, err := config.ConnectBucket(context.Background())

			if err != nil {
				panic(err)
			}

			// Set Address
			address := ip + ":" + port

			// Initalize Resources
			areaRepository := area_respository.NewRepository(db)
			areaUseCase := area_usecase.NewUseCase(areaRepository)
			areaController := controllers.NewAreaController(areaUseCase)

			assetsUseCase := assets_usecase.NewUseCase(minioClient)

			poiRepository := poi_respository.NewRepository(db)
			poiUseCase := poi_usecase.NewUseCase(poiRepository)
			poiController := controllers.NewPoiController(poiUseCase)

			queryAgentConnection := config.GPTService()
			queryAgentInitialize := query_agent_usecase.NewPromptServiceClient(queryAgentConnection)
			queryAgentUseCase := query_agent_usecase.NewUseCase(queryAgentInitialize)

			propertyRepository := property_respository.NewRepository(db)
			propertyUseCase := property_usecase.NewUseCase(propertyRepository, areaUseCase, *queryAgentUseCase)
			propertyController := controllers.NewPropertyController(propertyUseCase, assetsUseCase)

			searchRepository := search_respository.NewRepository(db)
			searchUseCase := search_usecase.NewUseCase(searchRepository)
			searchController := controllers.NewSearchController(searchUseCase)

			userRepository := users_respository.NewRepository(db)
			userUseCase := users_usecase.NewUseCase(userRepository)

			authController := controllers.NewAuthController(userUseCase)

			// Initialize Controllers
			controllers := &routes.Controllers{
				AreaController:     areaController,
				AuthController:     authController,
				SearchController:   searchController,
				PropertyController: propertyController,
				PoiController:      poiController,
			}

			app := fiber.New()

			// Middleware Logger
			app.Use(logger.New(logger.Config{
				Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
			}))

			// Register Routes
			routes.RegisterRoutes(controllers, app)

			// Run Server
			if err := app.Listen(address); err != nil {
				panic(err)
			}
		},
	}
)
