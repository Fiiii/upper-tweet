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

	settings := postgresql.ConnectionURL{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK TEST V2",
		})
	})

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

	//for i, book := range books {
	//	log.Printf("Book %d: %#v\n", i, book)
	//}

	r.Run("127.0.0.1:8080")
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
