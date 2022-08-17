# Server

**TODO: Add description**

## Installation

If [available in Hex](https://hex.pm/docs/publish), the package can be installed
by adding `server` to your list of dependencies in `mix.exs`:

```elixir
def deps do
  [
    {:server, "~> 0.1.0"}
  ]
end
```

Documentation can be generated with [ExDoc](https://github.com/elixir-lang/ex_doc)
and published on [HexDocs](https://hexdocs.pm). Once published, the docs can
be found at [https://hexdocs.pm/server](https://hexdocs.pm/server).



protoc for golang: 
protoc -I ../proto -I ../proto/include --elixir_out=plugins=grpc:lib/api --plugin=/root/.mix/escripts/protoc-gen-elixir ../proto/server.proto  

protoc for elixir:
protoc -I ../proto -I ../proto/include --elixir_out=plugins=grpc:lib/api --plugin=/Users/amirhasanbasic/.asdf/installs/elixir/1.12.3-otp-24/.mix/escripts/protoc-gen-elixir ../proto/server.proto 