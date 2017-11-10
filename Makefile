all: source

.PHONY: clean
clean:
	rm ttts

.PHONY: source
source:
	go build -o ttt-server
