package main

import (
	"fmt"
	"sync"

	"github.com/aleroxac/goexpert-events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			go rabbitmq.Publish(ch, fmt.Sprintf("msg: %d", i), "amq.direct")
		}
		wg.Done()
	}()
	wg.Wait()
}
