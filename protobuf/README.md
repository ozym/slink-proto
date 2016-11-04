go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

protoc --go_out=plugins=grpc:. slinkpb.proto
