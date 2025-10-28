package main

import (
	"log"

	"github.com/duonghuu/go_starter_01/users"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=example dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	users.Migrate(db)

	repo := users.NewRepository(db)
	handler := users.NewHandler(repo)

	r := gin.Default()
	handler.RegisterRoutes(r)

	log.Println("Server is running on http://localhost:3000")
	r.Run(":3000")
}
