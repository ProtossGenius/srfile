package svr_rpc_fileitf

//Product by SureMoonNet
//Author: ProtossGenius
//Auto-code should not change.
import "github.com/ProtossGenius/srfile/fileitf"
import "github.com/golang/protobuf/proto"
import "github.com/ProtossGenius/srfile/pb/smn_dict"
import "github.com/ProtossGenius/SureMoonNet/pb/smn_base"
import "net"
import "fmt"
import "github.com/ProtossGenius/srfile/pb/rip_fileitf"

type SvrRpcFileReqItf struct {
	itf   fileitf.FileReqItf
	dicts []smn_dict.EDict
}

func NewSvrRpcFileReqItf(itf fileitf.FileReqItf) *SvrRpcFileReqItf {
	list := make([]smn_dict.EDict, 0)
	list = append(list, smn_dict.EDict_rip_fileitf_FileReqItf_CreateFile_Prm)
	list = append(list, smn_dict.EDict_rip_fileitf_FileReqItf_OpenFile_Prm)
	list = append(list, smn_dict.EDict_rip_fileitf_FileReqItf_Read_Prm)
	list = append(list, smn_dict.EDict_rip_fileitf_FileReqItf_Write_Prm)
	list = append(list, smn_dict.EDict_rip_fileitf_FileReqItf_MakeSureAccessCode_Prm)
	return &SvrRpcFileReqItf{itf: itf, dicts: list}
}
func (this *SvrRpcFileReqItf) getEDictList() []smn_dict.EDict {
	return this.dicts
}
func ReadEdict_rip_fileitf_FileReqItf_CreateFile_Prm(bytes []byte) *rip_fileitf.FileReqItf_CreateFile_Prm {
	msg := &rip_fileitf.FileReqItf_CreateFile_Prm{}
	err := proto.Unmarshal(bytes, msg)
	if err != nil {
		panic(err)
	}
	return msg
}
func ReadEdict_rip_fileitf_FileReqItf_OpenFile_Prm(bytes []byte) *rip_fileitf.FileReqItf_OpenFile_Prm {
	msg := &rip_fileitf.FileReqItf_OpenFile_Prm{}
	err := proto.Unmarshal(bytes, msg)
	if err != nil {
		panic(err)
	}
	return msg
}
func ReadEdict_rip_fileitf_FileReqItf_Read_Prm(bytes []byte) *rip_fileitf.FileReqItf_Read_Prm {
	msg := &rip_fileitf.FileReqItf_Read_Prm{}
	err := proto.Unmarshal(bytes, msg)
	if err != nil {
		panic(err)
	}
	return msg
}
func ReadEdict_rip_fileitf_FileReqItf_Write_Prm(bytes []byte) *rip_fileitf.FileReqItf_Write_Prm {
	msg := &rip_fileitf.FileReqItf_Write_Prm{}
	err := proto.Unmarshal(bytes, msg)
	if err != nil {
		panic(err)
	}
	return msg
}
func ReadEdict_rip_fileitf_FileReqItf_MakeSureAccessCode_Prm(bytes []byte) *rip_fileitf.FileReqItf_MakeSureAccessCode_Prm {
	msg := &rip_fileitf.FileReqItf_MakeSureAccessCode_Prm{}
	err := proto.Unmarshal(bytes, msg)
	if err != nil {
		panic(err)
	}
	return msg
}
func (this *SvrRpcFileReqItf) OnMessage(c *smn_base.Call, conn net.Conn) (_d int32, _p proto.Message, _e error) {
	defer func() {
		if err := recover(); err != nil {
			_p = nil
			_e = fmt.Errorf("%v", err)
		}
	}()
	switch smn_dict.EDict(c.Dict) {
	case smn_dict.EDict_rip_fileitf_FileReqItf_CreateFile_Prm:
		{
			_msg := ReadEdict_rip_fileitf_FileReqItf_CreateFile_Prm(c.Msg)
			_d = int32(smn_dict.EDict_rip_fileitf_FileReqItf_CreateFile_Ret)
			p0, p1 := this.itf.CreateFile(_msg.Hash, _msg.Size)
			return _d, &rip_fileitf.FileReqItf_CreateFile_Ret{FileCode: p0, Err: p1}, nil
		}
	case smn_dict.EDict_rip_fileitf_FileReqItf_OpenFile_Prm:
		{
			_msg := ReadEdict_rip_fileitf_FileReqItf_OpenFile_Prm(c.Msg)
			_d = int32(smn_dict.EDict_rip_fileitf_FileReqItf_OpenFile_Ret)
			p0, p1, p2 := this.itf.OpenFile(_msg.Hash)
			return _d, &rip_fileitf.FileReqItf_OpenFile_Ret{FileCode: p0, Size: p1, Err: p2}, nil
		}
	case smn_dict.EDict_rip_fileitf_FileReqItf_Read_Prm:
		{
			_msg := ReadEdict_rip_fileitf_FileReqItf_Read_Prm(c.Msg)
			_d = int32(smn_dict.EDict_rip_fileitf_FileReqItf_Read_Ret)
			p0, p1 := this.itf.Read(_msg.FileCode, _msg.Block)
			return _d, &rip_fileitf.FileReqItf_Read_Ret{AccessCode: p0, Err: p1}, nil
		}
	case smn_dict.EDict_rip_fileitf_FileReqItf_Write_Prm:
		{
			_msg := ReadEdict_rip_fileitf_FileReqItf_Write_Prm(c.Msg)
			_d = int32(smn_dict.EDict_rip_fileitf_FileReqItf_Write_Ret)
			p0, p1 := this.itf.Write(_msg.FileCode, _msg.Block)
			return _d, &rip_fileitf.FileReqItf_Write_Ret{AccessCode: p0, Err: p1}, nil
		}
	case smn_dict.EDict_rip_fileitf_FileReqItf_MakeSureAccessCode_Prm:
		{
			_msg := ReadEdict_rip_fileitf_FileReqItf_MakeSureAccessCode_Prm(c.Msg)
			_d = int32(smn_dict.EDict_rip_fileitf_FileReqItf_MakeSureAccessCode_Ret)
			this.itf.MakeSureAccessCode(_msg.AccessCode)
			return _d, &rip_fileitf.FileReqItf_MakeSureAccessCode_Ret{}, nil
		}
	}
	return -1, nil, nil
}
