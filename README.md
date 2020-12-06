# GOSERV
server template in GO

## Run Server
go run cmd/go-server/main.go

## generating go-server code
using swagger-codegen 3.0.23

### Go API Server for swagger
```
java -jar ./api/swagger-codegen-cli.jar generate  -i ./api/swagger-server.yaml -l go-server -o ./cmd/go-server
```

### Go API client for swagger
```
java -jar ./api/swagger-codegen-cli.jar generate  -i ./api/swagger-client.yaml -l go -o ./internal/clients/sample
```

### DB Schema for ent
```
go run github.com/facebook/ent/cmd/ent init ${your-new-schema} 
```

edit ${your-new-schema}.   
initially User,Project Schema is set :)

```
go generate ./ent
```
