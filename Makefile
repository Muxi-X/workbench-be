build:
	go build -o main
proto:
	protoc --go_out=plugins=micro:. ./proto/status.proto