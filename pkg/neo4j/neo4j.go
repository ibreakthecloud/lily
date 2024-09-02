package neo4j

import (
	"context"
	"log"

	"github.com/ibreakthecloud/lily/pkg/models"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	Neo4jURI      = "bolt://neo4j:7687"
	Neo4jUsername = "neo4j"
	Neo4jPassword = "password"
)

type Neo4jClient struct {
	Driver neo4j.DriverWithContext
}

func NewNeo4jClient() (neo4j.DriverWithContext, error) {
	return neo4j.NewDriverWithContext(Neo4jURI, neo4j.BasicAuth(Neo4jUsername, Neo4jPassword, ""))
}

// Store annotation in Neo4j
func StoreAnnotationInNeo4j(ctx context.Context, annotation models.Annotation, driver neo4j.DriverWithContext) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// Create or match the entity node
	_, err := session.Run(ctx, "MERGE (e:Entity {name: $entity_name}) "+
		"MERGE (a:Annotation {type: $type, description: $description}) "+
		"MERGE (e)-[:ANNOTATED_AS]->(a)",
		map[string]interface{}{
			"entity_name": annotation.EntityName,
			"type":        annotation.Type,
			"description": annotation.Description,
		})
	if err != nil {
		log.Printf("Failed to store annotation in Neo4j: %s", err)
	}
}
