client:
	go run ./cmd/client/main.go

server:
	go run ./cmd/server/main.go

build:
	GOOS=windows GOARCH=amd64 go build -o release/windows/typex-client.exe -ldflags "-s -w" ./cmd/client/main.go
	GOOS=windows GOARCH=amd64 go build -o release/windows/typex-server.exe -ldflags "-s -w" ./cmd/server/main.go
	go build -o release/linux/typex-client -ldflags "-s -w" ./cmd/client/main.go
	go build -o release/linux/typex-server -ldflags "-s -w" ./cmd/server/main.go