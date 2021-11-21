Player = 3
Name = Hoge

client:
	go run ./cmd/client/main.go -name ${Name}

server:
	go run ./cmd/server/main.go -player ${Player}

build:
	GOOS=windows GOARCH=amd64 go build -o release/windows/typex-client.exe -ldflags "-s -w" ./cmd/client/main.go
	GOOS=windows GOARCH=amd64 go build -o release/windows/typex-server.exe -ldflags "-s -w" ./cmd/server/main.go
	go build -o release/linux/typex-client -ldflags "-s -w" ./cmd/client/main.go
	go build -o release/linux/typex-server -ldflags "-s -w" ./cmd/server/main.go