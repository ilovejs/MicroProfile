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

		addr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable",
			"postgres",
			"123123",
			"postgres",
		)

		//addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable",
		//  cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

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
		es, err := event.NewNats(fmt.Sprintf("nats://%s", "0.0.0.0:4222"))
		//es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))

		if err != nil {
			log.Println(err)
			return err
		}
		event.SetEventStore(es)
		return nil
	})
	defer event.Close()

  r := mux.NewRouter()
  r.HandleFunc("/profiles", createProfileHandler).Methods("POST", "OPTIONS")
  r.HandleFunc("/profiles/{id:[0-9]+}", UpdateProfileHandler).Methods("PUT", "OPTIONS")

  log.Print("Profile-Services")

  handler := cors.Default().Handler(r)
  if err := http.ListenAndServe(":8081", handler); err != nil {
    log.Fatal(err)
  }
}
