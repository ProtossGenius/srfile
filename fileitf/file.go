package fileitf

// FileReqItf file interface.
// it only accept command and not trans file. it's expect short block.
// error usually is a json, give new server or another solution.
type FileReqItf interface {
	// CreateFile create a new file.
	CreateFile(fileHash string, fileSize int64) (fileCode int64, err string)
	// OpenFile @path is file's virtual(or net) path.
	OpenFile(fileHash string) (fileCode, fileSize int64, err string)
	// Read read from fileCode.
	Read(fileCode, block int64) (accessCode int32, blockSize int64, preHash, subHash, err string)
	// Write write to fileCode.
	Write(fileCode, block, blockSize int64, preHash, subHash string) (accessCode int32, err string)
	// MakeSureAccessCode tell server you are ready to deal accessCode.
	MakeSureAccessCode(accessCode int32)
}
