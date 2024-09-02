package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ibreakthecloud/lily/pkg/models"
)

type LilyProducer struct {
	Producer *kafka.Producer
}

var Producer *LilyProducer

func NewProducer() (*LilyProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "broker:9092"})
	if err != nil {
		return nil, err
	}
	return &LilyProducer{Producer: p}, nil
}

func (p *LilyProducer) Close() {
	p.Producer.Close()
}

// Produce data annotation to Kafka
func (p *LilyProducer) ProduceDataAnnotation(annotation models.Annotation) {
	topic := "atlan_data_annotations"
	annotationJSON, _ := json.Marshal(annotation)
	p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          annotationJSON,
	}, nil)
}

// Produce Monte Carlo data issues to Kafka
func (p *LilyProducer) ProduceMonteCarlo(incident models.DataIssue) {
	topic := "monte_carlo_data_issues"

	incidentJSON, _ := json.Marshal(incident)
	p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          incidentJSON,
	}, nil)
}
