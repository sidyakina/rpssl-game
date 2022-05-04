# rpssl-game

This is a system with two microservices for the game "Rock Paper Scissors Spock Lizard" (http://www.samkass.com/theories/RPSSL.html)

## Credentials:
* github.com/caarlos0/env/v6 - library for parsing env variables
* github.com/gogo/protobuf - library for proto files
* github.com/pkg/errors - library for errors
* github.com/stretchr/testify - library for tests
* google.golang.org/grpc - library for grpc

## Starting system:
make build-services
make start-services

## Stopping system:
make stop-services

## Additional:
I have added a field "message" to "/play" that contains an extended result for users.