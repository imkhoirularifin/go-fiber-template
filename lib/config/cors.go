package config

import "github.com/gofiber/fiber/v2/middleware/cors"

var CorsConfig = cors.Config{
	AllowOrigins:     "http://localhost:3000",
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	AllowHeaders:     "*",
	AllowCredentials: true,
}
