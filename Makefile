test:
	CGO_ENABLED=1 go test -p 1 -race -cover -count=1 -coverprofile=.coverprofile -coverpkg=./... ./internal/...

protoc:
	protoc internal/service/grpc/pb/library.proto --go_out=./ --go-grpc_out=./
