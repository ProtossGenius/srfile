package clt_rpc_fileitf

//Product by SureMoonNet
//Author: ProtossGenius
//Auto-code should not change.
import "github.com/ProtossGenius/srfile/fileitf"
import "github.com/golang/protobuf/proto"
import "github.com/ProtossGenius/srfile/pb/rip_fileitf"
import "github.com/ProtossGenius/srfile/pb/smn_dict"
import "github.com/ProtossGenius/SureMoonNet/smn/net_libs/smn_rpc"
import "sync"

type CltRpcFileReqItf struct {
	fileitf.FileReqItf
	conn smn_rpc.MessageAdapterItf
	lock sync.Mutex
}

func NewCltRpcFileReqItf(conn smn_rpc.MessageAdapterItf) *CltRpcFileReqItf {
	return &CltRpcFileReqItf{conn: conn}
}
func (this *CltRpcFileReqItf) CreateFile(hash string, size int64) (int64, string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_msg := &rip_fileitf.FileReqItf_CreateFile_Prm{Hash: hash, Size: size}
	this.conn.WriteCall(int32(smn_dict.EDict_rip_fileitf_FileReqItf_CreateFile_Prm), _msg)
	_rm, _err := this.conn.ReadRet()
	if _err != nil {
		panic(_err)
	}
	if _rm.Err {
		panic(string(_rm.Msg))
	}
	_res := &rip_fileitf.FileReqItf_CreateFile_Ret{}
	_err = proto.Unmarshal(_rm.Msg, _res)
	if _err != nil {
		panic(_err)
	}
	return _res.FileCode, _res.Err
}
func (this *CltRpcFileReqItf) OpenFile(hash string) (int64, int64, string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_msg := &rip_fileitf.FileReqItf_OpenFile_Prm{Hash: hash}
	this.conn.WriteCall(int32(smn_dict.EDict_rip_fileitf_FileReqItf_OpenFile_Prm), _msg)
	_rm, _err := this.conn.ReadRet()
	if _err != nil {
		panic(_err)
	}
	if _rm.Err {
		panic(string(_rm.Msg))
	}
	_res := &rip_fileitf.FileReqItf_OpenFile_Ret{}
	_err = proto.Unmarshal(_rm.Msg, _res)
	if _err != nil {
		panic(_err)
	}
	return _res.FileCode, _res.Size, _res.Err
}
func (this *CltRpcFileReqItf) Read(fileCode int64, block int64) (int32, string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_msg := &rip_fileitf.FileReqItf_Read_Prm{FileCode: fileCode, Block: block}
	this.conn.WriteCall(int32(smn_dict.EDict_rip_fileitf_FileReqItf_Read_Prm), _msg)
	_rm, _err := this.conn.ReadRet()
	if _err != nil {
		panic(_err)
	}
	if _rm.Err {
		panic(string(_rm.Msg))
	}
	_res := &rip_fileitf.FileReqItf_Read_Ret{}
	_err = proto.Unmarshal(_rm.Msg, _res)
	if _err != nil {
		panic(_err)
	}
	return _res.AccessCode, _res.Err
}
func (this *CltRpcFileReqItf) Write(fileCode int64, block int64) (int32, string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_msg := &rip_fileitf.FileReqItf_Write_Prm{FileCode: fileCode, Block: block}
	this.conn.WriteCall(int32(smn_dict.EDict_rip_fileitf_FileReqItf_Write_Prm), _msg)
	_rm, _err := this.conn.ReadRet()
	if _err != nil {
		panic(_err)
	}
	if _rm.Err {
		panic(string(_rm.Msg))
	}
	_res := &rip_fileitf.FileReqItf_Write_Ret{}
	_err = proto.Unmarshal(_rm.Msg, _res)
	if _err != nil {
		panic(_err)
	}
	return _res.AccessCode, _res.Err
}
func (this *CltRpcFileReqItf) MakeSureAccessCode(accessCode int32) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_msg := &rip_fileitf.FileReqItf_MakeSureAccessCode_Prm{AccessCode: accessCode}
	this.conn.WriteCall(int32(smn_dict.EDict_rip_fileitf_FileReqItf_MakeSureAccessCode_Prm), _msg)
	_rm, _err := this.conn.ReadRet()
	if _err != nil {
		panic(_err)
	}
	if _rm.Err {
		panic(string(_rm.Msg))
	}
	_res := &rip_fileitf.FileReqItf_MakeSureAccessCode_Ret{}
	_err = proto.Unmarshal(_rm.Msg, _res)
	if _err != nil {
		panic(_err)
	}
	return
}
