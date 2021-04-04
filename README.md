# chainsafe-ethereum

This is a sample project which uses IPFS system and store the `cid` on ethereum blockchain. 

Note: 
      The project contains all the necessary files along the vendor files, due to compatibility issues in go-ethereum.
     
     
Prerequisites:
---------------

1. Go installed, latest version
2. Docker installed on your machine
3. IPFS installed on your machine. you can find it here: <a>https://docs.ipfs.io/install/command-line/#system-requirements</a>
4. The `make` tool is installed.

Setup:
--------------

1) Go to `/cmd/main` directory
2) run `make all` 
3) Copy the private key of an account out of any 10 accounts, displayed in terminal
4) update the `privatekey` in `env.yml` file, located in `cmd/env.yml`

5) <h5>Deploy Contract</h5> 

- Go to `cmd/contract`
- run `go run deploy.go`
- copy the contract address displayed in terminal after the deployment and update the `contractaddress` in the same `env.yml` file.

6) go to `cmd/main`
7) run `go run main.go sample1.jpeg` to store the image on ipfs and the resulting cid will be stored in enthereum blockchain.

Note: You can also change the file to the desired name in `go run main.go sample1.jpeg` command. `sample1.jpeg` is the sample file added already for you. 
      Make sure you add the file at `/cmd/main` location
