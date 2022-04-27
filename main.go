package main

import (
	"harmony/src/constants"
	"harmony/src/router"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", constants.GetDBInfo())

	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	harmonyRouter := gin.Default()
	router.SetupRouter(harmonyRouter, db)

	harmonyRouter.Run()
}
