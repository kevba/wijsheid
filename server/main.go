package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var dbPath = "./wisdom.db"

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}

	SetupDb(db)

	r.GET("/api/wisdoms", func(c *gin.Context) {
		wisdoms := GetWisdoms(db)
		c.JSON(http.StatusOK, gin.H{"wisdoms": wisdoms})
	})

	r.PUT("/api/wisdoms", func(c *gin.Context) {
		w := &Wisdom{}
		err := c.BindJSON(w)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}

		CreateWisdom(db, *w)
		c.Status(http.StatusCreated)
	})

	_ = r.Run("localhost:8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
