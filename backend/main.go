package main

import (
	"log"
	"os"
	"portfolio/apps"
	"portfolio/handlers"

	"github.com/go-playground/validator/v10"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	db := apps.ConnectToDatabase()
	validate := validator.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/api/v1/auth/sign-up" || c.Path() == "/api/v1/auth/sign-in" {
			return c.Next()
		}
		return jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECREET_KEY"))},
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.JSON(fiber.Map{"error": err.Error()})
			},
			SuccessHandler: func(c *fiber.Ctx) error {
				return c.Next()
			},
		})(c)
	})

	// auth
	apiAuth := app.Group("/api/v1/auth")
	apiAuth.Post("/sign-up", func(c *fiber.Ctx) error {
		return handlers.NewAuthHandlerCreate(c, db, validate)
	})
	apiAuth.Post("/sign-in", func(c *fiber.Ctx) error {
		return handlers.NewAuthHandlerSignIn(c, db, validate)
	})
	apiAuth.Delete("/delete-acount/:id", func(c *fiber.Ctx) error {
		return handlers.NewAuthHandlerDelete(c, db)
	})

	// project
	apiProj := app.Group("/api/v1/project")
	apiProj.Post("/create-project", func(c *fiber.Ctx) error {
		return handlers.NewProjectHandlerCreate(c, db, validate)
	})
	apiProj.Put("/update-project/:id", func(c *fiber.Ctx) error {
		return handlers.NewProjectHandlerUpdate(c, db, validate)
	})
	apiProj.Delete("/delete-project/:id", func(c *fiber.Ctx) error {
		return handlers.NewProjectHandlerDelete(c, db)
	})
	apiProj.Get("/get-project", func(c *fiber.Ctx) error {
		return handlers.NewProjectHandlerGetAllProject(c, db)
	})

	// skill
	apiSkill := app.Group("/api/v1/skill")
	apiSkill.Post("/create-skill", func(c *fiber.Ctx) error {
		return handlers.NewSkillHandlerCreate(c, db, validate)
	})
	apiSkill.Put("/update-skill/:id", func(c *fiber.Ctx) error {
		return handlers.NewSkillHandlerUpdate(c, db, validate)
	})
	apiSkill.Delete("/delete-skill/:id", func(c *fiber.Ctx) error {
		return handlers.NewSkillHandlerDelete(c, db)
	})
	apiSkill.Get("/get-skill", func(c *fiber.Ctx) error {
		return handlers.NewSkillHandlerGetAll(c, db)
	})

	log.Fatal(app.Listen(":8000"))
}
