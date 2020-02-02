# Copyright 2020 Vladislav Smirnov

all: build test

dep_install:
	go get github.com/gorilla/mux
	go get github.com/spf13/viper

build: dep_install
	go build

test:
	go test -v