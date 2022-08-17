#!/usr/bin/env sh
protoc -I /home/protoc/source -I /home/protoc/source/include --elixir_out=plugins=grpc:/home/protoc/code/lib/internal_api --plugin=/root/.mix/escripts/protoc-gen-elixir /home/protoc/source/server.proto


protoc -I ../proto -I ../proto/include --elixir_out=plugins=grpc:lib/api --plugin=/Users/amirhasanbasic/.asdf/installs/elixir/1.12.3-otp-24/.mix/escripts/protoc-gen-elixir ../proto/server.proto  