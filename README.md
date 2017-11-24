# Goethapi

Work in progress

## Swagger
### Validate API
```
swagger validate config/swagger.yml
```
The swagger spec at "config/swagger.yml" is valid against swagger specification 2.0

### Generate API
```
swagger generate server -A Goethapi -f ./config/swagger.yml
```

## Smart contract

### Compile contract
Execute
```
abigen --sol=contracts/token.sol --pkg=contracts --out=contracts/token.go
```

Run
```
go run src/github.com/charonne/goethapi/cmd/goethapi-server/main.go --host 0.0.0.0 --port 8080
```
