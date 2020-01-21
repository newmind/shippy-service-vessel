build:
	protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto
	docker build -t shippy-service-vessel .

run:
	docker run --rm -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns shippy-service-vessel