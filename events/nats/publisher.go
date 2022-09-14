package nats

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
)

type Publisher struct {
	conn    *stan.Conn
	subject string
}

func NewPublisher(conn *stan.Conn, subject string) *Publisher {
	return &Publisher{
		conn:    conn,
		subject: subject,
	}
}

func (p *Publisher) Publish(event any) error {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = (*p.conn).Publish(p.subject, eventJson)
	if err != nil {
		return err
	}

	return nil
}
