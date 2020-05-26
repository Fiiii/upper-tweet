package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"upper.io/db.v3/postgresql"
)

type Book struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	AuthorID  int    `db:"author_id"`
	SubjectID int    `db:"subject_id"`
}

func main() {
	initEnvs()
	InitServer().Run(os.Getenv("API_PORT"))

	settings := postgresql.ConnectionURL{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	sees, err := postgresql.Open(settings)
	if err != nil {
		log.Fatalf("db.open() fakap: %q\n", err)
	}
	defer sees.Close()

	//var books []Book
	//err = sees.Collection("Books").Find().All(&books)
	//if err != nil {
	//	log.Fatalf("Find(): %q\n", err)
	//}
}

func InitServer() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", pingEndpoint)

	return r
}

func pingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func initEnvs() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}
