package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/recipes/fiber-grpc/proto"
	"golangtest.com/client/auth/routes"
	"golangtest.com/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	app := fiber.New()
	models.ConnectDatabase()
	routes.Routes(app)
	app.Use(logger.New())

	log.Fatal(app.Listen(":3000"))

}
