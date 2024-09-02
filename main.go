package main

import (
	"context"
	"log"

	kfka "github.com/ibreakthecloud/lily/pkg/kafka"
	neo "github.com/ibreakthecloud/lily/pkg/neo4j"
	"github.com/ibreakthecloud/lily/pkg/server"
)

func main() {
	var err error
	neo4jClient, err := neo.NewNeo4jClient()
	if err != nil {
		log.Fatalf("Failed to create Neo4j client: %s", err)
	}

	ctx := context.Background()

	defer neo4jClient.Close(ctx)

	kfka.InitTopics(ctx) // Initialize Kafka topics

	go kfka.ConsumeDataIssues(ctx, neo4jClient)      // Start consuming inbound data issues
	go kfka.ConsumeDataAnnotations(ctx, neo4jClient) // Start consuming outbound data annotations

	r := server.InitServer()
	r.Run(":8080")
}
