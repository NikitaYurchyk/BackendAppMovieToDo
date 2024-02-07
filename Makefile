.PHONY:
.SILENT:


build:
	go build -o ./.bin/movie_to-do cmd/main.go
run: build
	./.bin/movie_to-do