# client
cd ./client
rm -rf src/lib/proto
mkdir src/lib/proto
npm run proto

# server
cd ../server
rm -rf proto
mkdir proto
protoc --go_out=./proto --go_opt=paths=source_relative \
    --go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
    --proto_path=../proto \
    ../proto/*.proto

