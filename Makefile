.PHONY: lint

TAGS ?=

GPB ?= 3.6.1
GPB_IMG ?= znly/protoc:0.4.0

COMPOSE ?= docker-compose
RUN ?= docker run --rm --user $$(id -u):$$(id -g)
PROTOC ?= $(RUN) -v "$$PWD:$$PWD" -w "$$PWD" $(GPB_IMG)
PROTOLOCK ?= $(RUN) -v $$PWD:/protolock -w /protolock nilslice/protolock

all: pkg/noterepl.pb.go lint
	go vet $(TAGS) ./...
	EXE=srv && cd cmd/$$EXE && CGO_ENABLED=0 go build -o $$EXE $(if $(wildcard $$EXE),|| rm $$EXE)

update: SHELL := /bin/bash
update:
	[[ 'libprotoc $(GPB)' = $$($(PROTOC) --version) ]]
	cd cmd/srv && go get -u -a -v ./...
	cd cmd/clt && go get -u -a -v ./...
	go mod tidy
	go mod verify

pkg/noterepl.pb.go: noterepl.proto
	$(PROTOC) -I=. --go_out=plugins=grpc:pkg $^

lint: SHELL = /bin/bash
lint:
	go fmt ./...
	$(PROTOLOCK) commit


debug: GRPC_HOST ?= localhost:7890
debug: all
	GRPC_HOST=$(GRPC_HOST) ./cmd/srv/srv

test: GRPC_HOST ?= localhost:7890
test: all
	grpcurl -proto noterepl.proto -d '{"language":"python3", "code":"print(42)\nprint(21*2)"}' -plaintext $(GRPC_HOST) NoteREPL/Eval
