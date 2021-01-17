# GOSERV
server template in GO

## Run API Server
go run cmd/goserv-api/main.go

## Run DB Server
go run cmd/goserv-db/main.go

## Run Rabbit MQ with Docker
```
docker run --rm -d --name rabbitmq -p 5672:5672 -p 15672:15672  -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin rabbitmq:3-management
```

## Generating Code
using openapi-generator 5.0.0-beta3

### API Server
```
java -jar ./api/openapi-generator-cli.jar generate -i ./api/swagger-server.yaml -g go-gin-server -o ./internal/goserv-api --package-name api
```

### API Client (sample)
```
java -jar ./api/openapi-generator-cli.jar generate -i ./api/swagger-client.yaml -g go -o ./internal/clients/sample --package-name client
```
Use ``api/codgen-server.go`` to create ``swagger-server.yaml``

### DB Schema for ent
```
go run github.com/facebook/ent/cmd/ent init ${your-new-schema} 
```

edit ``<your-new-schema>``. initially ``User``,``Project`` Schema is set :)

```
go generate ./ent
```
