package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ilovejs/profile/db"
	"github.com/ilovejs/profile/event"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"github.com/tinrab/retry"
)

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress      string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to PostgreSQL
	retry.ForeverSleep(2*time.Second, func(attempt int) error {
		// local
		//addr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable",
		//  "postgres", "123123", "postgres", )

		// docker-compose
		addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable",
			cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

		repo, err := db.NewPostgres(addr)
		if err != nil {
			log.Println(err)
			return err
		}
		db.SetRepository(repo)
		return nil
	})
	defer db.Close()

	// Connect to NATs
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		// local
		//es, err := event.NewNats(fmt.Sprintf("nats://%s", "0.0.0.0:4222"))
		// docker-compose
		es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))

		if err != nil {
			log.Println(err)
			return err
		}
		event.SetEventStore(es)
		return nil
	})
	defer event.Close()

	r := mux.NewRouter()

	// Attach "OPTIONS" for extra info
	r.HandleFunc("/profiles", listProfilesHandler).Methods("GET")
	r.HandleFunc("/profiles", createProfileHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/profiles/{id:[0-9]+}", UpdateProfileHandler).Methods("PUT", "OPTIONS")

	log.Print("Profile-Services")

	// Default() is not useful when PUT
	handler := cors.AllowAll().Handler(r)
	// change to 8081 for debugging
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
