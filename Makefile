.DEFAULT_GOAL := help

EXPLICIT := help test fmt vet build run tidy clean

.PHONY: $(EXPLICIT)

help:
	@echo "Usage:"
	@echo "  make <name>          Create new problem dir <name>/<name>.go"
	@echo "                       e.g. make 219_contains_duplicate"
	@echo "  make test            go test ./..."
	@echo "  make fmt             go fmt ./..."
	@echo "  make vet             go vet ./..."
	@echo "  make build           go build ./..."
	@echo "  make run DIR=<name>  go run ./<name>"
	@echo "  make tidy            go mod tidy"

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	go build ./...

tidy:
	go mod tidy

run:
	@test -n "$(DIR)" || (echo "DIR=<name> required" >&2; exit 1)
	go run ./$(DIR)

# Dynamically register any non-explicit goal as a phony target so that
# `make 219_contains_duplicate` works regardless of whether a directory
# with that name already exists.
NEW_TARGETS := $(filter-out $(EXPLICIT),$(MAKECMDGOALS))

ifneq ($(NEW_TARGETS),)
.PHONY: $(NEW_TARGETS)
$(NEW_TARGETS):
	@test ! -d "$@" || (echo "directory $@ already exists" >&2; exit 1)
	@mkdir "$@"
	@printf 'package main\n\nfunc main() {}\n' > "$@/$@.go"
	@echo "created $@/$@.go"
endif
