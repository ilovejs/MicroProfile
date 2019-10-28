package event

import "github.com/ilovejs/profile/schema"

type EventStore interface {
	Close()
	PublishProfileCreated(profile schema.Profile) error
	PublishProfileUpdated(profile schema.Profile) error
	OnProfileCreated(f func(ProfileCreatedMessage)) error
}

var impl EventStore

func SetEventStore(es EventStore) {
	impl = es
}

func Close() {
	impl.Close()
}

func PublishProfileCreated(p schema.Profile) error {
	return impl.PublishProfileCreated(p)
}

func PublishProfileUpdated(p schema.Profile) error {
	return impl.PublishProfileUpdated(p)
}

func OnProfileCreated(f func(ProfileCreatedMessage)) error {
	return impl.OnProfileCreated(f)
}
