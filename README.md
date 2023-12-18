# protocolbuf

# compile
protoc version=libprotoc 3.17.3
protoc --go_out=. --go-grpc_out=. gamefee_manager.proto
protoc --go_out=. --go-grpc_out=. skytree_backend.proto
