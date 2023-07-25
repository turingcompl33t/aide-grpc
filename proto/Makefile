# Defines builds for AIDE API

# The path to protocol buffers input source for aide
aide_src_path := src/aide
# The path to protocol buffers input source for google
google_src_path := src/google

# Glob that captures a list of all aide source files
aide_source_files := $(shell find $(aide_src_path) -type f -name '*.proto')
# Glob that captures a list of all google source files
google_source_files := $(shell find $(google_src_path) -type f -name '*.proto')

# TODO
grpc_in := $(aide_src_path)/api/v1/api.proto
# TODO
swagger_in := $(aide_src_path)/api/v1/api.proto

# The root build path
build_root := build

# The path to the protocol buffers build
build_proto := $(build_root)/proto

# The path to the build output directory for golang
build_proto_go := $(build_proto)/go

# The path to the build output directory for swagger
build_swagger := $(build_root)/swagger

all: build

build: build/proto/go build/swagger

.PHONY: build/proto/go
build/proto/go: $(aide_source_files)
	rm -rf $(build_proto_go) && mkdir -p $(build_proto_go)

	# Go protocol buffers generation
	for source in $(aide_source_files) ; do python -m grpc_tools.protoc -I src "$$source" --go_out=$(build_proto_go) || exit 1; done

	# gRPC generation
	python -m grpc_tools.protoc -I src $(grpc_in) \
		--go-grpc_out=$(build_proto_go) \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(build_proto_go) \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt allow_delete_body=true
	
	# Create the go package
	mv $(build_proto_go)/aide/api/v1/* $(build_proto_go)/aide/proto/pkg/apiv1
	mv $(build_proto_go)/aide/proto/pkg pkg

.PHONY: build/swagger
build/swagger: $(aide_source_files)
	rm -rf $(build_swagger) && mkdir -p $(build_swagger)
	
	python -m grpc_tools.protoc -I src $(swagger_in) \
		--swagger_out=logtostderr=true,allow_delete_body=true,json_names_for_fields=true:$(build_swagger)

.PHONY: clean
clean:
	rm -rf $(build_root) && rm -rf pkg
