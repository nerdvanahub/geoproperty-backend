package routes

import (
	"geoproperty_be/delivery/http/controllers"
	"geoproperty_be/domain"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mbndr/figlet4go"
)

type Controllers struct {
	AreaController     *controllers.AreaController
	AuthController     *controllers.AuthController
	SearchController   *controllers.SearchController
	PropertyController *controllers.PropertyController
	PoiController      *controllers.PoiController
}

func RegisterRoutes(c *Controllers, ctx *fiber.App) {
	// Root Route for Health Check
	ctx.Get("/", func(c *fiber.Ctx) error {
		// Figlet Banner
		ascci := figlet4go.NewAsciiRender()
		banner, _ := ascci.Render("GeoProperty Backend")

		return c.SendString(banner)
	})

	// Docs Route API
	ctx.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("https://documenter.getpostman.com/view/11577849/2s9YXh532M")
	})

	// Add Middleware CORS
	ctx.Use(cors.New())

	// Compresed Middleware
	ctx.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// API Version 1
	v1 := ctx.Group("/api/v1")

	// Register Area Routes
	areaRoutes := v1.Group("/area")
	areaRoutes.Post("/overlaps", c.AreaController.CheckOverlaps)

	// Register Auth Routes
	authRoutes := v1.Group("/auth")
	authRoutes.Get("/callback", c.AuthController.CallbackGoogle)
	authRoutes.Get("/login/google", c.AuthController.LoginGoogle)
	authRoutes.Post("/login", c.AuthController.Login)
	authRoutes.Post("/register", c.AuthController.Register)

	// Register Search Routes
	searchRoutes := v1.Group("/search")
	searchRoutes.Get("/:keyword", c.SearchController.Search)
	searchRoutes.Post("/all", c.SearchController.GetAll)

	// Register Poi Routes
	poiRoutes := v1.Group("/poi")
	poiRoutes.Post("/", c.PoiController.GeoJSON)

	// Register Property Routes
	propertyRoutes := v1.Group("/property")
	propertyRoutes.Post("/point", c.PropertyController.GetByCenterPoint)
	propertyRoutes.Post("/prompt", c.PropertyController.GetByPrompt)
	propertyRoutes.Get("/:uid", c.PropertyController.GetDetail)

	// Middleware JWT
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				Status:  fiber.StatusUnauthorized,
				Message: "Unauthorized",
			})
		},
	}))

	// Restrict Property Routes
	propertyRoutes.Post("/", c.PropertyController.Insert)
	propertyRoutes.Post("/own", c.PropertyController.GetOwn)
	propertyRoutes.Delete("/:uid", c.PropertyController.Delete)
}
