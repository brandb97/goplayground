client: client.go main.go
	go build main.go client.go -o client
server: pipeserver.go main.go
	go build main.go pipeserver.go -o server
clean:
	rm server client
