build-services:
	docker build -t game-service -f ./build/game-service/Dockerfile .
	docker build -t gateway -f ./build/gateway/Dockerfile .

start-services:
	/bin/bash ./run-services.sh

stop-services:
	docker rm -f gateway game-service
