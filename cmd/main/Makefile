all: run-docker run-ipfs

build-docker:
	 docker build -t ganache .

run-docker: build-docker
	docker run -p 8545:8545 ganache

run-ipfs:
		$$(ipfs daemon)