proto:
	protoc --proto_path=proto/ --go_out=. proto/*

clean:
	rm pb/*.go

run:
	go run main.go

test:
	go test -race -timeout 10s -coverprofile=/tmp/cover.out ./...

.PHONY: proto clean run test
