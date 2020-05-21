test:

clean:

build:
	go get -u github.com/ProtossGenius/smntools/cmd/smnrpc-autocode
	smnrpc-autocode -cfg ./datas/cfgs/rpcfile.json

