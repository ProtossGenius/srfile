package fileitf

//FileReqItf file interface.
//it only accept command and not trans file. it's expect short block.
//error usually is a json, give new server or another solution
type FileReqItf interface {
	//OpenFile @path is file's virtual(or net) path.
	OpenFile(path string) (fileCode int64, err string)
	//Read read from fileCode. start is file.byte[start], size is how many byte to read.
	Read(fileCode, start, size int64) (accessCode int32, err string)
	//Write write to fileCode. if start not end, will overwrite, so be careful.
	Write(fileCode, start, size int64) (accessCode int32, err string)
	//MakeSureAccessCode tell server you are ready to deal accessCode.
	MakeSureAccessCode(accessCode int32)
}
