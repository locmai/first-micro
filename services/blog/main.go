package main

import (
	"os"

	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
	"github.com/locmai/first-micro/services/blog/handler"
	"github.com/locmai/first-micro/services/blog/model"
	"github.com/locmai/first-micro/services/blog/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	blog "github.com/locmai/first-micro/services/blog/proto/blog"
)

func main() {
	// Initialize environment
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.blog"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	db := pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASS"),
		Addr:     os.Getenv("POSTGRES_HOST") + ":" + os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
	})

	defer db.Close()

	err = model.CreateSchema(db)

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to database!")

	// Register Handler
	blog.RegisterBlogHandler(service.Server(), &handler.Blog{Db: db})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.blog", service.Server(), new(subscriber.Blog))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.blog", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
