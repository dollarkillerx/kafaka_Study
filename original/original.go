package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)


	test1()
}

func test1() {
	//producer()
	consumer()
}

func producer() {
	topic := "test-top"
	partition := 0

	toCtx, _ := context.WithTimeout(context.Background(), time.Second*3)
	conn, err := kafka.DialLeader(toCtx, "tcp", "192.168.88.11:9091", topic, partition)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	log.Println("av")
	//conn.SetWriteDeadline(time.Now().Add(10 * time.Second)) // 超过时间 停止写入
	for i := 0; i < 100; i++ {
		messages, err := conn.WriteMessages(
			kafka.Message{Value: []byte(fmt.Sprintf("Hello %d", i))},
		)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(messages)
	}

	log.Println("Send Success")
}

func consumer() {
	log.Println("Consumer Init")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"192.168.88.11:9091"},
		GroupID:  "sc02",
		Topic:    "assets",
		MinBytes: 10e3,
		MaxBytes: 1e6,
	})

	for {
		message, err := reader.ReadMessage(context.TODO())
		if err != nil {
			log.Println(err)
		}

		log.Println("key: ", string(message.Key))
		fmt.Println("val: ", string(message.Value))
		break
	}
}