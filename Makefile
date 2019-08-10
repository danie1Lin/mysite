GO = go


.PHONY: all
all: help

.PHONY: api
api:  
	go run api/main.go start

.PHONY: update_pb
update_pb:
	protoc -I=pkg/user/pb --go_out=plugins=grpc:pkg/user/pb pkg/user/pb/user.proto

.PHONY: setup
setup: 
	go get -u github.com/golang/protobuf/protoc-gen-go;
	go get -u google.golang.org/grpc;

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'