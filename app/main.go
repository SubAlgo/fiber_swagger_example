// ไฟล์ app/main.go
package main

import (
	"log"

	// import folder learnFiberSwagger/doc แต่หากกำหนดชื่อ module เป็นชื่ออื่นก็ให้แก้จาก learnFiberSwagger เป็นชื่อที่เราได้ตั้งไว้

	"learnFiberSwagger/app/routes"
	_ "learnFiberSwagger/doc"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title User API by Fiber and Swagger
// @version 1.0
// @description API user management Server by Fiber | Doc by Swagger.

// @contact.name admin
// @contact.url http://subalgo.com/support
// @contact.email admin@subalgo.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Fiber instance
	app := fiber.New()

	app.Get("/swagger/*", swagger.Handler)

	routes.RegisterRoute(app)
	// Start Server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
