package neo4j

import (
	"context"
	"fmt"
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

	// todo: fix this
	query := `
		MERGE (e:%s {id: $entity_name, name: $entity_name})
		MERGE (n:Annotation {type: $type})
		MERGE (e)-[:ANNOTATED_AS]->(n)
		SET n.description = $description
	`

	query = fmt.Sprintf(query, annotation.EntityType)
	_, err := session.Run(ctx, query,
		map[string]interface{}{
			"entity_name": annotation.EntityName,
			"entity_type": annotation.EntityType,
			"type":        annotation.Type,
			"description": annotation.Description,
		})
	if err != nil {
		log.Printf("Failed to store annotation in Neo4j: %s", err)
	}
}

func StoreIssueInNeo4j(ctx context.Context, issue models.DataIssue, driver neo4j.DriverWithContext) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	issueId := issue.TableName + issue.ColumnName + issue.IssueType

	_, err := session.Run(ctx, `
				MERGE (t:Table {name: $table_name, id: $table_name})
				MERGE (c:Column {name: $column_name, id: $column_name})
				MERGE (t)-[:HAS_COLUMN]->(c)
				CREATE (i:DataIssue {id: $issue_id, table_name: $table_name, column_name: $column_name, issue_type: $issue_type, issue_severity: $issue_severity})
				MERGE (c)-[:HAS_ISSUE]->(i)
				`,
		map[string]interface{}{
			"table_name":     issue.TableName,
			"column_name":    issue.ColumnName,
			"issue_type":     issue.IssueType,
			"issue_severity": issue.IssueSeverity,
			"issue_id":       issueId,
		})
	if err != nil {
		log.Printf("Failed to insert into Neo4j: %s", err)
	}
}
