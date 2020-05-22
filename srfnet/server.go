package srfnet

import (
	"net"

	"github.com/ProtossGenius/SureMoonNet/basis/smn_net"
	"github.com/ProtossGenius/SureMoonNet/smn/net_libs/muti_service"
	"github.com/ProtossGenius/SureMoonNet/smn/net_libs/smn_rpc"
	"github.com/ProtossGenius/srfile/rpcfile/svr_rpc_fileitf"
)

//SRFileServer suremoon rpc file server..
type SRFileServer struct {
	fileMgr *SRFileMgr
	svcMgr  *muti_service.ServiceManager
}

//OpenFile @path is file's virtual(or net) path.
func (s *SRFileServer) OpenFile(path string) (fileCode int64, err string) {
	code, e := s.fileMgr.OpenFile(path)
	return code, e.Error()
}

//Read read from fileCode. start is file.byte[start], size is how many byte to read.
func (s *SRFileServer) Read(fileCode int64, start int64, size int64) (accessCode int32, err string) {
	panic("not implemented") // TODO: Implement
}

//Write write to fileCode. if start not end, will overwrite, so be careful.
func (s *SRFileServer) Write(fileCode int64, start int64, size int64) (accessCode int32, err string) {
	panic("not implemented") // TODO: Implement
}

//MakeSureAccessCode tell server you are ready to deal accessCode.
func (s *SRFileServer) MakeSureAccessCode(accessCode int32) {
	panic("not implemented") // TODO: Implement
}

func accept() func(conn net.Conn) {
	mgr := NewSRFileMgr()

	return func(conn net.Conn) {
		sm := muti_service.NewServiceManager(conn)
		svr := svr_rpc_fileitf.NewSvrRpcFileReqItf(&SRFileServer{fileMgr: mgr, svcMgr: sm})
		smn_rpc.ServiceManagerRegister(sm, 1, "smrfile", svr)
	}
}

//RunServer run srfile server.
func RunServer(port int) error {
	doAccept := accept()

	svr, err := smn_net.NewTcpServer(port, doAccept)
	if err != nil {
		return err
	}

	svr.Run()

	return nil
}
