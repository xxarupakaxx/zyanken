.PHONY: proto
proto:
	@protoc -I . proto/*.proto --go_out=gen/pb --go_opt=paths=source_relative --go-grpc_out=gen/pb --go-grpc_opt=paths=source_relative