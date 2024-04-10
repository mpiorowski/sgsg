# Client
rm -rf ./client/src/lib/proto
mkdir ./client/src/lib/proto
proto-loader-gen-types --keepCase --longs=String --enums=Number --defaults --oneofs --grpcLib=@grpc/grpc-js --outDir=./client/src/lib/proto ./proto/*.proto && cp ./proto/*.proto ./client/src/lib/proto/

# Service Auth
rm -rf ./service-auth/proto
mkdir ./service-auth/proto
protoc --go_out=./service-auth/proto --go_opt=paths=source_relative \
    --go-grpc_out=./service-auth/proto --go-grpc_opt=paths=source_relative \
    --proto_path=./proto \
    ./proto/*.proto

# Service Profile
rm -rf ./service-profile/proto
mkdir ./service-profile/proto
protoc --go_out=./service-profile/proto --go_opt=paths=source_relative \
    --go-grpc_out=./service-profile/proto --go-grpc_opt=paths=source_relative \
    --proto_path=./proto \
    ./proto/*.proto
