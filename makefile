.PHONY: tailwind-watch
tailwind-watch:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate -watch -proxy=http://localhost:3000 

.PHONY: dev
dev:
	go build -o ./tmp/ && air serve

.PHONY: build
build:
	make tailwind-build && make templ-generate && go build main.go
