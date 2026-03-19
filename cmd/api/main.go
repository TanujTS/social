package main

import (
	"log"
	"os"

	"github.com/tanujts/social/internal/env"
	"github.com/tanujts/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	os.LookupEnv("PATH")

	mux := app.mount()
	log.Fatal(app.run(mux))
}
