package nats

import (
	"github.com/nats-io/stan.go"
)

// Config is the required properties to use the nats client.
type Config struct {
	ClusterId string
	ClientId  string
}

// Connect knows how to establish a connection with the streaming service.
func Connect(config Config) (*stan.Conn, error) {
	sc, err := stan.Connect(config.ClusterId, config.ClientId)
	if err != nil {
		return nil, err
	}

	return &sc, nil
}
