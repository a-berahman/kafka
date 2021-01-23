package consumer

import (
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

//First ...
func First(ctx context.Context) {
	//create consumer config
	topic := "tpc-first"
	partition := 2
	connectionType := "tcp"
	connectionAddress := "localhost:9092"
	//create and subscribe consumer to our topics
	connection, err := kafka.DialLeader(ctx, connectionType, connectionAddress, topic, partition)
	defer connection.Close()
	if err != nil {
		log.Fatal(err)
	}

	connection.SetReadDeadline((time.Now().Add(10 * time.Second)))
	//pull for new data
	readBatch := connection.ReadBatch(500, 500000)
	defer readBatch.Close()
	byteString := make([]byte, 500)
	for {
		_, err := readBatch.Read(byteString)
		if err != nil {
			break
		}
		fmt.Printf("message value: %v\n", string(byteString))
	}
}

//Second ...
func Second(ctx context.Context) {
	//create consumer config
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "mobile-app",
		Topic:   "tpc-first",
	})

	//pull for new data
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("consume at topic: %v partition: %v offset: %v value: %s \n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal(err)
	}
}
