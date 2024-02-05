PLUGIN_ID ?= com.example.my-go-plugin
PLUGIN_VERSION ?= 0.1.0
BUNDLE_NAME ?= $(PLUGIN_ID)-$(PLUGIN_VERSION).tar.gz

.PHONY: server
server:
	mkdir -p ./server/dist
	GOOS=linux GOARCH=amd64 go build -o ./server/dist/plugin-linux-amd64 ./server
	tar -czvf ./server/dist/$(BUNDLE_NAME) ./server/dist/plugin-linux-amd64 plugin.json;
