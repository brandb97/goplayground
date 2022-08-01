client: main.go client.go pipeserver.go
	go build main.go client.go pipeserver.go
	mv main client
server: main.go pipeserver.go
	go build main.go client.go pipeserver.go
	mv main server
clean:
	rm server client
