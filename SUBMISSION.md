### Language
- [Go](https://golang.org/ "Go")

### Preparing environment on Windows

1. Install Go v1.12.5 (https://golang.org/dl/)
2. Clone this repository on path: **%GOPATH%/src**
3. Open command window and go to cloned path: **cd %GOPATH%/src/back-end-take-home**
3. Download dependencies with the command: **go get**
5. Run the application with the command: **go run %GOPATH%/src/back-end-take-home/main.go**

### Get shortest route

> Method: GET
Endpoint: /flight

#### Example

##### Request:
> GET http://localhost:8080/flight?origin=Toronto,Canada&destination=Rio de Janeiro,Brazil

##### Response:
```json
[
    {
        "airline": {
            "code": "UAL",
            "name": "United Airlines",
            "country": "United States"
        },
        "origin": {
            "name": "Lester B. Pearson International Airport",
            "city": "Toronto",
            "country": "Canada",
            "iata": "YYZ",
            "latitude": "43.67720032",
            "longitude": "-79.63059998"
        },
        "destination": {
            "name": "George Bush Intercontinental Houston Airport",
            "city": "Houston",
            "country": "United States",
            "iata": "IAH",
            "latitude": "29.9843998",
            "longitude": "-95.34140015"
        }
    },
    {
        "airline": {
            "code": "UAL",
            "name": "United Airlines",
            "country": "United States"
        },
        "origin": {
            "name": "George Bush Intercontinental Houston Airport",
            "city": "Houston",
            "country": "United States",
            "iata": "IAH",
            "latitude": "29.9843998",
            "longitude": "-95.34140015"
        },
        "destination": {
            "name": "Rio Galeão – Tom Jobim International Airport",
            "city": "Rio De Janeiro",
            "country": "Brazil",
            "iata": "GIG",
            "latitude": "-22.80999947",
            "longitude": "-43.25055695"
        }
    }
]
```
