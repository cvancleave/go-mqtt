# go-mqtt

Contains a simple client for connecting, subscribing, and publishing via mqtt.

### Running Locally

Setup a local broker via docker.
- `docker run -d --name emqx -p 1883:1883 emqx:latest`

Run clientA to connect and subscribe to a topic.
- `go run cmd/clientA/main.go`

Run clientB to connect and publish to a topic.
- `go run cmd/clientA/main.go`