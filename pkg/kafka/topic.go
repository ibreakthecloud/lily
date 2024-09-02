package kafka

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	MonteCarloTopic     = "monte_carlo_data_issues"
	DataAnnotationTopic = "atlan_data_annotations"
)

func InitTopics(ctx context.Context) {
	// check if topic exists
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "broker:9092"})
	if err != nil {
		log.Fatalf("Failed to create admin client: %s", err)
	}
	defer admin.Close()

	// Create topics if they don't exist
	topicSpecs := []kafka.TopicSpecification{
		{
			Topic:             MonteCarloTopic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
		{
			Topic:             DataAnnotationTopic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	for _, topicSpec := range topicSpecs {
		_, err = admin.CreateTopics(ctx, []kafka.TopicSpecification{topicSpec})
		if err != nil {
			log.Printf("Failed to create topic %s: %s", topicSpec.Topic, err)
		}
	}

}
