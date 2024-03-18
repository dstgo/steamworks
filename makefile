target_proto_files := $(shell find ./protobufs -name *.proto)

.PHONY: gen_pb
gen_pb:
	protoc --proto_path=./protobufs \
		   --go_out=paths=source_relative:./message \
		   $(target_proto_files)