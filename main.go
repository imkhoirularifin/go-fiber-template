package main

import "go-fiber-template/internal/infrastructure"

//go:generate swag init --parseDependency

//	@title			Go Fiber Template API Documentation
//	@version		1.0
//	@description	Go Fiber Template API Documentation

//	@BasePath	/api/v1
//	@schemes	http https

//	@accept		application/json
//	@produce	application/json

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	infrastructure.Run()
}
