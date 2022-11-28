package main

import (
	"RabbitMq-Demo/delayletter/constant"
	"RabbitMq-Demo/util"
	"log"
)

func main() {
	mq := util.NewRabbitMQ()
	defer mq.Close()
	mqCh := mq.Channel

	msgsCh, err := mqCh.Consume(
		constant.Queue1,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	util.FailOnError(err, "消费队列失败")
	forever := make(chan bool)
	go func() {
		for d := range msgsCh {
			log.Printf("接收到的消息: %s", d.Body)
			d.Ack(false)
		}
	}()
	log.Printf("[*] Waiting for message, To exit press CTRL+C")
	<-forever
}
