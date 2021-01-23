package main

import (
	"context"

	"github.com/a-berahman/kafka/app/consumer"
	"github.com/a-berahman/kafka/app/producer"
)

func main() {
	ctx := context.Background()
	go producer.Second(ctx)
	consumer.Second(ctx)
}
