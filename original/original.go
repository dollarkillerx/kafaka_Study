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
	producer()
	consumer()
}

func producer() {
	topic := "test-top"
	partition := 0

	toCtx, _ := context.WithTimeout(context.Background(), time.Second*3)
	conn, err := kafka.DialLeader(toCtx, "tcp", "192.168.88.11:9292", topic, partition)
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
	topic := "test-top"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "192.168.88.11:21811", topic, partition)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//conn.SetReadDeadline(time.Now().Add(10*time.Second))  设置读取timeout

	// 设置读取批次 (最小粒度,最大粒度)
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	defer batch.Close()
	b := make([]byte,10e3)
	for {
		read, err := batch.Read(b)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(string(read))
	}
}