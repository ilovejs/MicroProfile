package event

import (
	"bytes"
	"encoding/gob"

	"github.com/ilovejs/profile/schema"
	"github.com/nats-io/go-nats"
)

type NatsEventStore struct {
	nc                         *nats.Conn
	profileCreatedSubscription *nats.Subscription
	profileCreatedChan         chan ProfileCreatedMessage
}

func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

// callback
func (e *NatsEventStore) OnProfileCreated(f func(ProfileCreatedMessage)) (err error) {
	m := ProfileCreatedMessage{}
	e.profileCreatedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		// read message data into m
		_ = e.readMessage(msg.Data, &m)
		f(m)
	})
	return
}

func (e *NatsEventStore) Close() {
	if e.nc != nil {
		e.nc.Close()
	}
	if e.profileCreatedSubscription != nil {
		_ = e.profileCreatedSubscription.Unsubscribe()
	}
	close(e.profileCreatedChan)
}

//create event
func (e *NatsEventStore) PublishProfileCreated(profile schema.Profile) error {
	m := ProfileCreatedMessage{profile.ID, profile.Name}
	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	// publish
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventStore) PublishProfileUpdated(profile schema.Profile) error {
	m := ProfileUpdatedMessage{profile.ID, profile.Name}
	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	// publish
	return e.nc.Publish(m.Key(), data)
}

// utils
func (e *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (e *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	// append to buffer
	b.Write(data)
	// store in m
	return gob.NewDecoder(&b).Decode(m)
}
