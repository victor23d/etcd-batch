-include .env

VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID := /tmp/.$(PROJECTNAME).pid


test:
	go clean -testcache
	docker stop etcd || true
	docker run -d --rm --name etcd -p 2379:2379 \
	-e ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
	-e ETCD_ADVERTISE_CLIENT_URLS=http://0.0.0.0:2379 \
	quay.io/coreos/etcd:latest \
	# etcd \
	# --listen-client-urls http://0.0.0.0:2379 \
	# --advertise-client-urls http://0.0.0.0:2379 
	# vendor broken
	# go test -v -mod=vendor ./...
	go test -v ./...
	docker stop etcd


all: help test
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
