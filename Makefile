all: compile strip

compile:
	go build

strip:
	strip gotag

install:
	install -Dm755 ./gotag /usr/bin/gotag
