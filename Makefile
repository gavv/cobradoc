SHELL := bash

all: tidy build

build:
	go build .
	cd _example && go build -o example

example: build
	./_example/example markdown > _example/_markdown.md
	( echo '```' ; man <(./_example/example man) ; echo '```' ) > _example/_manpage.md

tidy:
	go mod tidy -v
	cd _example && go mod tidy -v

fmt:
	gofmt -s -w . ./_example

md:
	md-authors --format modern --append AUTHORS.md
