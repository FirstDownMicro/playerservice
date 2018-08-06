buildlocal:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/Nanixel/FirstDownMicro/playerservice \
		proto/playerservice.proto
build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/Nanixel/FirstDownMicro/playerservice \
		proto/playerservice.proto
	docker build -t playerservice .
runlocal:
	go run *.go --server_address :50053 --registry mdns
run:
	docker run --net=host \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e MICRO_REGISTRY=mdns \
		playerservice
		#-e DB_HOST=localhost \
		-e DB_USER=jessetellez \
		-e DB_NAME=FootballDB \
		-e DB_PASSWORD=password1