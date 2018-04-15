SHELL := /bin/bash

PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
SERVE_PORT := 8181

build:
	@docker build \
	-t my_test \
	.

build-clean:
	@docker build \
	--no-cache \
	-t my_test \
	.

run-dev:
	docker run \
	--rm \
	-v $(PROJECT_DIR):/go/src/app \
	-p $(SERVE_PORT):8181 \
	my_test

build-and-run: build-clean
build-and-run: run-dev
