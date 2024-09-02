package kafka

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ibreakthecloud/lily/pkg/models"
)

// Produce data annotation to Kafka
func ProduceDataAnnotation(annotation models.Annotation) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "broker:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer p.Close()

	topic := "atlan_data_annotations"
	annotationJSON, _ := json.Marshal(annotation)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          annotationJSON,
	}, nil)
	p.Flush(15 * 1000)
}

// Produce Monte Carlo data issues to Kafka
func ProduceMonteCarlo(incident models.DataIssue) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "broker:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer p.Close()

	topic := "monte_carlo_data_issues"

	incidentJSON, _ := json.Marshal(incident)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          incidentJSON,
	}, nil)
	p.Flush(15 * 1000)
}
