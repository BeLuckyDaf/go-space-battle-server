all: build run

build:
	@go build
	@echo Done.

run:
	@./go-space-battle-server.exe