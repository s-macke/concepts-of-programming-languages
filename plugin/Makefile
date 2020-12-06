.PHONY: help plugin master run

help:
	@echo "plugin - build the plugin"
	@echo "   run - build plugin and run master"

plugin:
	go build -buildmode=plugin plugin.go

run: plugin
	go run master.go plugin.so
