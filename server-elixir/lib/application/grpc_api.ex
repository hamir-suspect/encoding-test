defmodule Server.GrpcApi do
  use GRPC.Server, service: Api.Server.Server.Service

  require Logger

  def create(req, _) do
    {:ok, priv_key} = ExPublicKey.load("./priv.pem")

    # {:ok, encrypted_value} = Base.decode64(req.encrypted_value)

    {:ok, decrypted_val} = ExPublicKey.decrypt_private(req.encrypted_value, priv_key)

    some_msg = Api.Server.SomeMsg.decode(decrypted_val)

    Logger.info("some msg decoded: #{inspect some_msg}")

    encoded_msg = Api.Server.SomeMsg.encode(some_msg)

    {:ok, pub_key} = ExPublicKey.load("./pub.pem")

    {:ok, encrypted_msg} = ExPublicKey.encrypt_public(encoded_msg, pub_key)

    # encoded_msg = Base.encode64(encrypted_msg)

    Api.Server.CreateResponse.new(encrypted_value: encrypted_msg)
  end

  def create_native(req, _) do
    Api.Server.SomeMsg.new()
  end
end
