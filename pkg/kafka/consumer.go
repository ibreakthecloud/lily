package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/ibreakthecloud/lily/pkg/models"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	neo "github.com/ibreakthecloud/lily/pkg/neo4j"
)

// Consume data issues from Kafka
func ConsumeDataIssues(ctx context.Context, driver neo4j.DriverWithContext) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:9092",
		"group.id":          "data-issue-consumer",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{"monte_carlo_data_issues"}, nil)

	// Store in Neo4j
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var issue models.DataIssue
			json.Unmarshal(msg.Value, &issue)
			neo.StoreIssueInNeo4j(ctx, issue, driver)
		} else {
			log.Printf("Error while consuming message: %s", err)
		}
	}
}

// Consume data annotations from Kafka
func ConsumeDataAnnotations(ctx context.Context, driver neo4j.DriverWithContext) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:9092",
		"group.id":          "data-annotation-consumer",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{"atlan_data_annotations"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var annotation models.Annotation
			json.Unmarshal(msg.Value, &annotation)
			log.Printf("Received annotation: %s", string(msg.Value))
			neo.StoreAnnotationInNeo4j(ctx, annotation, driver)
			// all the downstream processing can be done here
		} else {
			log.Printf("Error while consuming annotation message: %s", err)
		}
	}
}
