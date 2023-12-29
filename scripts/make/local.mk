provision-local: ## provision local environment
	if [ ! -f db/todo.db ]; then \
		mkdir -p db \
		&& touch db/todo.db; \
	fi; \
	sqlite3 db/todo.db < db/schema.sql

build: ## build the service
	templ generate \
	&& cd cmd/reporting \
	&& go build -o bin/reporting

css: ## build minified css
	tailwindcss -i $(ROOT_PATH)/internal/assets/input.css -o $(ROOT_PATH)/internal/assets/output.css --minify

watch-css: ## watch for css changes and rebuild css
	tailwindcss -i $(ROOT_PATH)/internal/assets/input.css -o $(ROOT_PATH)/internal/assets/output.css --watchrun

.PHONY: build-css watch-css