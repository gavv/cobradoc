SHELL := bash

all:
	go build -o example
	./example markdown > _markdown.md
	( echo '```' ; man <(./example man) ; echo '```' ) > _manpage.md
