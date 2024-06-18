run-client:
	go run api/client/main.go
run-server:
	air -c .air.toml
run:
	make run-client & make run-server
proto:
	protoc --proto_path=api/proto --go_out=api/pb --go_opt=paths=source_relative \
	--go-grpc_out=api/pb --go-grpc_opt=paths=source_relative \
	api/proto/*.proto