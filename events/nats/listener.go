package nats

import (
	"github.com/nats-io/stan.go"
)

type Listener struct {
	conn        *stan.Conn
	msgHandler  stan.MsgHandler
	subject     string
	queueGroup  string
	durableName string
}

func NewListener(conn *stan.Conn, msgHandler stan.MsgHandler, subject, queueGroup, durableName string) *Listener {
	return &Listener{
		conn:        conn,
		msgHandler:  msgHandler,
		subject:     subject,
		queueGroup:  queueGroup,
		durableName: durableName,
	}
}

func (l *Listener) ListenQueue() (stan.Subscription, error) {
	sub, err := (*l.conn).QueueSubscribe(l.subject, l.queueGroup, l.msgHandler, stan.SetManualAckMode(),
		stan.DeliverAllAvailable(), stan.DurableName(l.durableName))
	if err != nil {
		return nil, err
	}

	return sub, nil
}
