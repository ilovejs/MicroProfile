package main

import (
	"fmt"
  "github.com/gorilla/handlers"
  "log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ilovejs/profile/db"
	"github.com/ilovejs/profile/event"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress      string `envconfig:"NATS_ADDRESS"`
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/profiles", listProfilesHandler).
		Methods("GET")
	return
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

		//addr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		//  "postgres", 5432, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

		repo, err := db.NewPostgres(addr)
		if err != nil {
			log.Println(err)
			return err
		}
		// Dependency injection
		db.SetRepository(repo)
		return nil
	})
	defer db.Close()

	// Connect to Nats
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		//es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
		es, err := event.NewNats(fmt.Sprintf("nats://%s", "0.0.0.0:4222"))
		if err != nil {
			log.Println(err)
			return err
		}

		// todo: handle nats event e.g. elastic search
		// or set es
		event.SetEventStore(es)
		return nil
	})
	defer event.Close()

	router := newRouter()
  headersOk := handlers.AllowedHeaders([]string{"X-Requested-With, Content-Type, Authorization"})
  originsOk := handlers.AllowedOrigins([]string{"*"}) // os.Getenv("ORIGIN_ALLOWED")
  methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
  if err := http.ListenAndServe(":8080",
    handlers.CORS(originsOk, headersOk, methodsOk)(router)); err != nil {
    log.Fatal(err)
  }
}
