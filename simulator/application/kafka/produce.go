package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route "github.com/rodrigocnascimento/code-delivery-simulator/application/route"
	"github.com/rodrigocnascimento/code-delivery-simulator/infra/kafka"
)

// Produce is responsible to publish the positions of each request
// Example of a json request:
//{"clientId":"1","routeId":"1"}
//{"clientId":"2","routeId":"2"}
//{"clientId":"3","routeId":"3"}
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	newRoute := route.NewRoute()
	json.Unmarshal(msg.Value, &newRoute)
	newRoute.LoadPositions()
	positions, err := newRoute.ExportJsonPositions()

	if err != nil {
		log.Println(err.Error())
	}
	
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}