package producer

import (
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

//First ...
func First(ctx context.Context) {
	//create producer properties
	topic := "tpc-first"
	partition := 2
	connectionType := "tcp"
	connectionAddress := "localhost:9092"

	//create the producer
	connection, err := kafka.DialLeader(ctx, connectionType, connectionAddress, topic, partition)
	defer connection.Close()
	if err != nil {
		log.Fatal(err)
	}
	connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
	//create producer message and send msg
	for i := 0; i < 10; i++ {
		connection.WriteMessages(
			kafka.Message{Value: []byte(fmt.Sprintf("message : %v", i))},
		)
	}
}

//Second ...
func Second(ctx context.Context) {
	//create producer properties
	topic := "tpc-first"
	//create the producer
	w := kafka.NewWriter(kafka.WriterConfig{Topic: topic, Brokers: []string{"localhost:9092"}})
	//create producer message and send message
	for i := 0; i < 10; i++ {
		err := w.WriteMessages(ctx, kafka.Message{Value: []byte(fmt.Sprintf("message from second producer - %v", i))})
		if err != nil {
			log.Fatal(err)
		}
	}
}
