package util

import (
	"github.com/streadway/amqp"
	"log"
)

const MqUrl = "amqp://guest:guest@127.0.0.1:5672/lxy"

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// NewRabbitMQ 拿到 rabbitmq 的 channel，轻量级 connection
func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp.Dial(MqUrl)
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	return &RabbitMQ{
		Conn: conn,
		Channel: ch,
	}
}

func (r RabbitMQ) Close()  {
	r.Channel.Close()
	r.Conn.Close()
}

// FailOnError 错误处理函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}