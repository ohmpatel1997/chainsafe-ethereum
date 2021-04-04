all: build run

build:
	 docker build -t ganache .

run: build
	docker run -p 8545:8545 ganache