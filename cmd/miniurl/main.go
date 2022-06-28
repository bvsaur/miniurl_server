package main

import (
	"github.com/bveranoc/mu_server/pkg/config"
	"github.com/bveranoc/mu_server/pkg/database"
	"github.com/bveranoc/mu_server/pkg/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	//* Middleware
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	//* Routes
	jwtConfig := middleware.JWTConfig{
		SigningKey:  []byte(config.GetEnv("JWT_SECRET")),
		TokenLookup: "header:Authorization",
	}

	m := e.Group("/minis")
	u := e.Group("/users")

	e.POST("/auth", handlers.Auth)
	m.Use(middleware.JWTWithConfig(jwtConfig))
	u.Use(middleware.JWTWithConfig(jwtConfig))

	u.PUT("", handlers.UpdateNickname)
	u.GET("/me", handlers.GetMe)

	e.GET("/minis/:mini", handlers.GetMini)
	m.POST("", handlers.CreateMini)
	m.GET("", handlers.GetMinis)
	m.DELETE("/:id", handlers.DeleteMini)

	database.ConnectDB()
	e.Logger.Fatal(e.Start(":8080"))
}
