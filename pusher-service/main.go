package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ilovejs/profile/event"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Nats
	hub := newHub()
	// try until success, fail to sleep
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		// event store
		es, err := event.NewNats(fmt.Sprintf("nats://%s", "0.0.0.0:4222"))
		//es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
		if err != nil {
			log.Println(err)
			return err
		}
		// Push messages to clients with callback
		err = es.OnProfileCreated(func(m event.ProfileCreatedMessage) {
			log.Printf("Profile received: %v\n", m)
			hub.broadcast(newProfileCreatedMessage(m.ID, m.Name), nil)
		})
		if err != nil {
			log.Println(err)
			return err
		}
		// set event store for Event
		event.SetEventStore(es)
		return nil
	})
	defer event.Close()

	// Run WebSocket server
	go hub.run()

	http.HandleFunc("/pusher", hub.handleWebSocket)

	// listen
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
