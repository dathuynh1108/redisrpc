proto:
	docker build -t protoc-builder ./rpc && \
	docker run -v $(CURDIR):/workspace protoc-builder \
	protoc \
	--go_opt=module=redisrpc --go_out=. \
	--go-grpc_opt=module=redisrpc --go-grpc_out=. \
	rpc/rpc.proto

