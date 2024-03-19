input := ./protobufs
output:= ./message
input_protos := $(shell find $(input) -name *.proto)

.PHONY: gen_pb
gen_pb:
	mkdir -p $(output)
	protoc --proto_path=./protobufs \
		   --go_out=paths=source_relative:$(output) \
		   $(input_protos)