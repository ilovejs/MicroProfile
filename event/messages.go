package event

type Message interface {
	Key() string
}

// create
type ProfileCreatedMessage struct {
	ID   string
	Name string
}

func (m *ProfileCreatedMessage) Key() string {
	return "profile.created"
}

// updates
type ProfileUpdatedMessage struct {
	ID   string
	Name string
}

func (m *ProfileUpdatedMessage) Key() string {
	return "profile.updated"
}
