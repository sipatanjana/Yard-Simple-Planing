package main

/*
Copyright © 2025 Dinas Komunikasi dan Informatika DIY <diskominfo@jogjaprov.go.id>
Pusat Layanan Transformasi Digital

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"YardPlaning/internal/conf"
	"YardPlaning/internal/middlewares"
	"YardPlaning/internal/repositories"
	"YardPlaning/internal/routes"
)

func main() {
	loadEnv()

	dsn := getDSN()

	db, err := repositories.InitDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	//init server
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Ganti dengan asal spesifik jika perlu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	server.Use(middlewares.DefaultHeaders())

	//testing server health
	server.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	//init routes
	container := conf.NewContainer(db)
	routes.InitRoutes(server, container)

	//start server
	server.Run(":" + os.Getenv("SERVER_PORT"))

	if err := server.Run(":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

/*
Helper function
*/

// load .env dari workspace directory
func loadEnv() {
	if err := godotenv.Load(filepath.Join("./", ".env")); err != nil {
		log.Fatal("Error load env")
	}
}

// assamble url database connection
func getDSN() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
}
