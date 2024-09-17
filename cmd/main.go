package main

import (
	"Building/internal/app"
)

// @title           Swagger Example API
// @version         1.0
// @description     Add/Get buildings.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  evilgooby1@gmail.com

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app.App()
}
