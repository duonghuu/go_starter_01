package main

import (
	"log"

	"github.com/duonghuu/go_starter_01/users"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var users = []User{
// 	{ID: 1, Name: "John", Email: "john@example.com"},
// 	{ID: 2, Name: "Alice", Email: "alice@example.com"},
// }

func main() {
	// connStr := "postgres://postgres:example@localhost:5432/postgres?sslmode=disable"

	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("Không thể kết nối tới database:", err)
	// }

	// r := gin.Default()

	// fmt.Println("✅ Kết nối thành công tới PostgreSQL!")

	// // CRUD routes
	// r.GET("/api/users", getUsers)
	// r.GET("/api/users/:id", getUserByID)
	// r.POST("/api/users", createUser)
	// r.PUT("/api/users/:id", updateUser)
	// r.DELETE("/api/users/:id", deleteUser)

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello from Gin!",
	// 	})
	// })

	// r.Run(":3000") // chạy ở http://localhost:3000

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

// GET /users
// func getUsers(c *gin.Context) {
// 	c.JSON(http.StatusOK, users)
// }

// // GET /users/:id
// func getUserByID(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	for _, u := range users {
// 		if u.ID == id {
// 			c.JSON(http.StatusOK, u)
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// }

// // POST /users
// func createUser(c *gin.Context) {
// 	var newUser User
// 	if err := c.ShouldBindJSON(&newUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	// Add new user
// 	newUser.ID = len(users) + 1
// 	users = append(users, newUser)
// 	c.JSON(http.StatusCreated, newUser)
// }

// // PUT /users/:id
// func updateUser(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	var updated User
// 	if err := c.ShouldBindJSON(&updated); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for i, u := range users {
// 		if u.ID == id {
// 			users[i].Name = updated.Name
// 			users[i].Email = updated.Email
// 			c.JSON(http.StatusOK, users[i])
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// }

// // DELETE /users/:id
// func deleteUser(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	for i, u := range users {
// 		if u.ID == id {
// 			users = append(users[:i], users[i+1:]...)
// 			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// }
