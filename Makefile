##Tail
debug:

qrun:

install:

test: build 
	 go run test/test_net/test.go

clean:
	rm -rf pb/ rpcfile/

build: clean
	#go get -u github.com/ProtossGenius/smntools/cmd/smnrpc-autocode # if not exist
	smnrpc-autocode -cfg ./datas/cfgs/rpcfile.json
	mv ./go/pb ./pb
	rm -rf ./go

