.PHONY: all server client clean

all: server client

server:
	cd server && go build -o ../server

client:
	cd client && go build -o ../client

clean:
	rm -f server client