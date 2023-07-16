PROTOF=$(shell find ./proto -iname "*.proto")

build:
	protoc -I ./proto --go_out=./lib/ --go_opt=paths=source_relative \
           --go-grpc_out=./lib/ --go-grpc_opt=paths=source_relative \
           ${PROTOF}
