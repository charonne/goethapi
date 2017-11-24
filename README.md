# FidEver

## Swagger
### Validate API
```
swagger validate config/swagger.yml
```
The swagger spec at "config/swagger.yml" is valid against swagger specification 2.0

### Generate API
```
swagger generate server -A FidEver -f ./config/swagger.yml
```





## Smart contract

### Compile contract
Execute
```
abigen --sol=contracts/token.sol --pkg=contracts --out=contracts/token.go
```






<hr>








Run
```
go run src/github.com/Phoax/gotestswagger/cmd/goethapi-server/main.go --host 0.0.0.0 --port 8080
```




Example

compile a contract
http://localhost:8080/contracts

{"name":"coucou", "source": "pragma solidity ^0.4.0;\n\n  contract SimpleStorage {\n      uint storedData;\n\n      function set(uint x) {\n          storedData = x;\n      }\n\n      function get() constant returns (uint) {\n          return storedData;\n      }\n  }"}
