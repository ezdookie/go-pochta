package main

import (
	"context"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/gogearbox/gearbox"
)

// DB Connection
var DB *pg.DB

func main() {
	// Connect to database
	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	DB = pg.Connect(opt)

	// Check connection is okay
	ctx := context.Background()
	if err := DB.Ping(ctx); err != nil {
		panic(err)
	}

	gb := gearbox.New()

	gb.Use(authMiddleware)
	gb.Post("/send", sendMessage)

	gb.Start(":8080")
}
