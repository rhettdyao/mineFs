package nameserver
type FileInfo struct{
	id uint64
	version uint64
	mode int
	blocks uint64
	ctime uint64
	name string
	size uint64
	parent_id uint64
	owner int
}

func (info *FileInfo)DecodeFileInfo(buf []byte) (bool){
	return false
}

func (info *FileInfo) EncodeFileInfo()([]byte){
	return nil
}
