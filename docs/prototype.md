# Prototype

The prototype for the Atlan Lily project will focus on implementing a simplified version of the metadata ingestion and consumption pipeline, using Golang as the primary programming language. The prototype will demonstrate the core functionality of metadata ingestion, storage, and consumption, without the full complexity of the authentication, authorization, and observability components.

## Design

1. User Authentication:
   -  User logs in and receives a JWT token.
   -  The token is used to authenticate subsequent requests.

2. Incident Reporting:
   - Authenticated users post data issues to the /monte-carlo endpoint.
   - Data issues are produced to the monte_carlo_data_issues Kafka topic.
   - Kafka consumer listens to this topic, processes messages, and stores them in Neo4j.

3. Data Annotation:
   - Authenticated users post annotations to the /annotate endpoint.
   - Annotations are produced to the atlan_data_annotations Kafka topic.
   - Kafka consumer listens to this topic, processes messages, and stores them in Neo4j.
   Note: In prototype data annotations is an api call to the server, in actual implementation it will be a part of the metadata transformation service.