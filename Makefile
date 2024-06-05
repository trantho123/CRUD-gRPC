run-client:
	go run api/client/main.go
run-server:
	air -c .air.toml
run:
	make run-client & make run-server
