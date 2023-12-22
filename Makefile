SHELL := bash
MAKE_PATH ?= $(abspath $(firstword $(MAKEFILE_LIST)))
ROOT_PATH ?= $(dir $(MAKE_PATH))

.DEFAULT: help
.PHONY: help

help: ## Print this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort \
	| sed 's/^.*\/\(.*\)/\1/' \
	| awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

include scripts/make/*.mk