defmodule Api.Server.SomeMsg.Code do
  @moduledoc false
  use Protobuf, enum: true, protoc_gen_elixir_version: "0.10.0", syntax: :proto3

  field :OK, 0
  field :FAILED, 2
end
defmodule Api.Server.SomeMsg do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.10.0", syntax: :proto3

  field :name, 1, type: :string
  field :code, 2, type: Api.Server.SomeMsg.Code, enum: true
  field :number, 3, type: :int32
  field :timestamp, 4, type: Google.Protobuf.Timestamp
end
defmodule Api.Server.CreateRequest do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.10.0", syntax: :proto3

  field :encrypted_value, 1, type: :string, json_name: "encryptedValue"
end
defmodule Api.Server.CreateResponse do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.10.0", syntax: :proto3

  field :encrypted_value, 1, type: :string, json_name: "encryptedValue"
end
defmodule Api.Server.Server.Service do
  @moduledoc false
  use GRPC.Service, name: "Api.Server.Server", protoc_gen_elixir_version: "0.10.0"

  rpc :Create, Api.Server.CreateRequest, Api.Server.CreateResponse

  rpc :CreateNative, Api.Server.SomeMsg, Api.Server.SomeMsg
end

defmodule Api.Server.Server.Stub do
  @moduledoc false
  use GRPC.Stub, service: Api.Server.Server.Service
end
