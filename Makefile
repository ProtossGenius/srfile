test: build 
	 go run test/test_net/test.go

clean:

build: build_clean
	go get -u github.com/ProtossGenius/smntools/cmd/smnrpc-autocode
	smnrpc-autocode -cfg ./datas/cfgs/rpcfile.json

build_clean:
	rm -rf ./rpcfile

