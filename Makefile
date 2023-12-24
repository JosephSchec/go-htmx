 
run:
	@templ generate
	@go build cmd/main.go
	@go run cmd/main.go