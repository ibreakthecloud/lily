# Atlan Lily

1. [Design](./docs/design.md)

2. [Prototype](./docs/prototype.md)

## Running the prototype

1. Make the docker image of the prototype
```bash
make build
```

2. Run the docker image
```bash
docker compose up -d
```

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Using the prototype

1. Login
```bash
curl -X POST http://localhost:8080/login -d '{"username": "admin", "password": "password"}'
```

2. Copy the JWT, we will use it for subsequent requests

3. Simulate monte-carlo events (data issue)
```bash
curl -X POST -H "Authorization: bearer <jwt-token>" localhost:8080/monte-carlo -d '{"table_name":"foo", "column_name":"foocol", "issue_type":"foo went wrong", "issue_severity":"high"}'
```

4. Simulate data annotation
```bash
curl -X POST -H "Authorization: bearer <jwt-token>" localhost:8080/annotate -d '{"entity_name": "foocol", "entity_type":"Column", "type":"PII", "description": "lorem ipsum"}'
```

5. Check neo4j for the data
    - Navigate to `http://localhost:7474/browser/`
    - Run the following querie
        - `MATCH (n)-[r]->(m) RETURN n, r, m`
    - You should see the data issues and annotations with the relationships
