package srfnet

import "io"

//NewSRFileMgr create.
func NewSRFileMgr() *SRFileMgr {
	res := &SRFileMgr{}

	return res
}

//SRFileMgr suremoon rpc file manager.
type SRFileMgr struct {
}

//WriteToNet read from local and write to net.
func (sfm *SRFileMgr) WriteToNet(w io.Writer, fileCode, start, size int64) error {
	return nil
}

//ReadFromNet read from net and write to local.
func (sfm *SRFileMgr) ReadFromNet(w io.Writer, fileCode, start, size int64) error {
	return nil
}

//OpenFile get path's fileCode.
func (sfm *SRFileMgr) OpenFile(path string) (fileCode int64, err error) {
	return 0, nil
}
