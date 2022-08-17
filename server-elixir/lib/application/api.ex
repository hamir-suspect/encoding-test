defmodule Server.Application do
  use Application
  require Logger

  def start(_type, _args) do
    # children = [{GRPC.Server.Supervisor, {[Server.GrpcApi, 50051]}}]
    children = [{GRPC.Server.Supervisor, {Server.GrpcApi, 50051}}]

    opts = [strategy: :one_for_one, name: Server.Supervisor]

    Logger.info("starting applications")

    Supervisor.start_link(children, opts)
  end
end
