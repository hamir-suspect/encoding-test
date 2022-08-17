
PROTOC_TAG=1.6.6-3.3.0-0.5.4

pb.gen.elixir:
	docker run --rm -v $(PWD)/server-elixir:/home/protoc/code -v proto:/home/protoc/source -t renderedtext/protoc:$(PROTOC_TAG) sh -c /home/protoc/code/scripts/internal_protos.sh
	scripts/vagrant_sudo chown -R $$(id -u $${USER}):$$(id -g $${USER}) lib/internal_api

pb.gen.golang:
	docker run --rm -v $(PWD)/client-go:/home/protoc/code -v proto:home/protoc/source -t renderedtext/protoc:$(PROTOC_TAG) sh -c /home/protoc/code/scripts/internal_protos.sh
	scripts/vagrant_sudo chown -R $$(id -u $${USER}):$$(id -g $${USER}) lib/internal_api