# server
rm -rf ./server/proto
mkdir ./server/proto
protoc --go_out=./server/proto --go_opt=paths=source_relative \
    --go-grpc_out=./server/proto --go-grpc_opt=paths=source_relative \
    --proto_path=./proto \
    ./proto/*.proto

# Client
rm -rf ./client/src/lib/proto
mkdir ./client/src/lib/proto
proto-loader-gen-types --keepCase --longs=String --enums=String --defaults --oneofs --grpcLib=@grpc/grpc-js --outDir=./client/src/lib/proto ./proto/*.proto && cp ./proto/*.proto ./client/src/lib/proto/
