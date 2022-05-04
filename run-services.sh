docker run --rm \
    --detach \
		--env-file game-service.env \
		--name game-service \
		--network host \
		game-service:latest

echo "waiting to start game-service"
sleep 30s

docker run --rm \
    --detach \
		--env-file gateway.env \
		--name gateway \
		--network host \
		gateway:latest
