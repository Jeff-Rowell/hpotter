package dockerclient

import "slices"

type OperationType int

const (
	Get OperationType = iota
	Append
	Delete
)

type ActorMessage struct {
	Type     OperationType
	Value    string
	Response chan []string
}

func manageContainers(messages <-chan ActorMessage) {
	containers := []string{}

	for message := range messages {
		switch message.Type {
		case Get:
			message.Response <- containers
		case Append:
			containers = append(containers, message.Value)
		case Delete:
			idx := slices.Index(containers, message.Value)
			if idx >= 0 {
				containers = append(containers[:idx], containers[idx+1:]...)
			}
		}
	}
}
