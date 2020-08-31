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

// Read read from fileCode.
func (r *RPCFile) Read(fileCode int64, block int64) (accessCode int32, size int64, preHash string, subHash string, err string) {
	panic("not implemented") // TODO: Implement
}

// Write write to fileCode.
func (r *RPCFile) Write(fileCode int64, block int64, size int64, preHash string, subHash string) (accessCode int32, err string) {
	panic("not implemented") // TODO: Implement
}

//CreateFile create a new file.
func (r *RPCFile) CreateFile(hash string, size int64) (fileCode int64, err string) {
	panic("not implemented") // TODO: Implement
}

//OpenFile @path is file's virtual(or net) path.
func (r *RPCFile) OpenFile(hash string) (fileCode int64, size int64, err string) {
	return -1, -1, "PANIC"
}

//MakeSureAccessCode tell server you are ready to deal accessCode.
func (r *RPCFile) MakeSureAccessCode(accessCode int32) {
	panic("not implemented") // TODO: Implement
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
		_, _ = adapter.WriteRet(dict, res, err)
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
