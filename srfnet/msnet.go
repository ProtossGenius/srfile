package srfnet

import "io"

const (
	//ErrFileNotFound when open file.
	ErrFileNotFound = "ErrFileNotFound"
	//ErrBlockExist should never rewrite a existed block.
	ErrBlockExist = "ErrBlockExist"
	//ErrInvalidBlock if block num < 0 or > max_block_id.
	ErrInvalidBlock = "ErrInvalidBlock"
	//ErrFileExist when create file.
	ErrFileExist = "ErrFileExist"
)

//NewSRFileMgr create.
func NewSRFileMgr() *SRFileMgr {
	res := &SRFileMgr{}

	return res
}

//SRFileMgr suremoon rpc file manager.
type SRFileMgr struct {
}

//GetReadCode When service want read from file. give access code.
func (sfm *SRFileMgr) GetReadCode(fileCode int64, start int64, size int64) (accessCode int32, err string) {
	return -1, ""
}

//GetWriteCode when service want write to file. give access code.
func (sfm *SRFileMgr) GetWriteCode(fileCode int64, start int64, size int64) (accessCode int32, err string) {
	return -1, ""
}

//WriteToNet read from local and write to net.
func (sfm *SRFileMgr) WriteToNet(w io.Writer, fileCode, start, size int64) error {
	return nil
}

//ReadFromNet read from net and write to local.
func (sfm *SRFileMgr) ReadFromNet(w io.Reader, fileCode, start, size int64) error {
	return nil
}

//OpenFile get path's fileCode.
func (sfm *SRFileMgr) OpenFile(path string) (fileCode int64, err error) {
	return 0, nil
}
