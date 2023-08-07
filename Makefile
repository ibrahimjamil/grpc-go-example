generate_grpc_code:
	protoc \
	--go_out=. \
	--go-grpc_out=. \
	proto/greet.proto