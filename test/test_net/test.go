package main

import (
	"fmt"
	"net"
	"time"

	"github.com/ProtossGenius/SureMoonNet/basis/smn_net"
	"github.com/ProtossGenius/SureMoonNet/smn/net_libs/smn_rpc"
	"github.com/ProtossGenius/srfile/rpcfile/clt_rpc_fileitf"
	"github.com/ProtossGenius/srfile/rpcfile/svr_rpc_fileitf"
)

const Port = 900

type RPCFile struct {
}

//OpenFile @path is file's virtual(or net) path.
func (r *RPCFile) OpenFile(path string) (fileCode int64, err string) {
	fmt.Println("In server path = ", path)
	return -1, "ERR_PANIC"
}

//Read read from fileCode. start is file.byte[start], size is how many byte to read.
func (r *RPCFile) Read(fileCode int64, start int64, size int64) (accessCode int32, err string) {
	panic("not implemented")
}

//Write write to fileCode. only support append.
func (r *RPCFile) Write(fileCode int64, size int64) (accessCode int32, err string) {
	panic("not implemented")
}

//MakeSureAccessCode tell server you are ready to deal accessCode.
func (r *RPCFile) MakeSureAccessCode(accessCode int32) {
	panic("not implemented")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func AccpterRun(adapter smn_rpc.MessageAdapterItf) {
	rpcSvr := svr_rpc_fileitf.NewSvrRpcFileReqItf(&RPCFile{})

	for {
		msg, err := adapter.ReadCall()
		check(err)
		dict, res, err := rpcSvr.OnMessage(msg, adapter.GetConn())

		adapter.WriteRet(int32(dict), res, err)
	}
}

func accept(conn net.Conn) {
	adapter := smn_rpc.NewMessageAdapter(conn)
	go AccpterRun(adapter)
}

func RunSvr() {
	svr, err := smn_net.NewTcpServer(Port, accept)
	check(err)
	svr.Run()
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	go RunSvr()
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", Port))
	check(err)

	client := clt_rpc_fileitf.NewCltRpcFileReqItf(smn_rpc.NewMessageAdapter(conn))
	fmt.Println(client.OpenFile("[param: path]"))
	client.MakeSureAccessCode(-1)
}
