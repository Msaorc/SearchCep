# SearchZipCode
Zip code search application

### Operation:
* **Search Cep**: To search for the zip code, execute the command via terminal: `go run main.go 00000-000`
* **Compile Application**: Execute the command via terminal: `go mod init namePackage`, `go build main.go` 
* **Execute linux**: `./main 00000-000`

# SearchZipCodeAPI
API zip code consultation

### Operation:
* **Search ZipCode API**: Execute the command via terminal: `go run server/main.go`

### EndPoint:
* **search**: `localhost:9000/search?zipCode=00000-000` 