# protocolbuf

# compile
protoc version=libprotoc 3.17.3
protoc --go_out=./skytree_backend/ --proto_path=./ skytree_backend.proto
