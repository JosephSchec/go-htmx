 
run: build
	@./tmp/

templ:
	@templ generate -watch -proxy=http://localhost:3000

tailwind:
	@tailwindcss -i view/css/app.css -o public/styles.css --watch

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss

build:
	tailwindcss -i view/css/app.css -o public/styles.css
	@templ generate view
	@go build -o ./tmp/main.exe ./cmd/main.go
