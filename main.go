package main

import (
	"harmony/src/constants"
	"harmony/src/router"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", constants.GetDBInfo())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	harmonyRouter := gin.Default()
	harmonyRouter.Use(cors.New(config))
	router.SetupRouter(harmonyRouter, db)

	harmonyRouter.Run()
}
