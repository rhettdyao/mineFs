package common

import (
	"os"
	"encoding/json"
)

type FileAttr struct {
	Name string "the file name"
	Mode os.FileMode
	Blocks int64 "the blocks"
	Ctime, Mtime, Atime int64
	Uid int32
	Gid int32
	StIno uint64
	ParentStIno uint64
}

func (attr *FileAttr) IsDir() bool{
	return attr.Mode & os.ModeDir != 0
}

func (attr *FileAttr) IsRegular() bool{
	return attr.Mode & os.ModeType == 0
}

func (attr FileAttr) Perm() os.FileMode{
	return attr.Mode & os.ModePerm
}

func (attr FileAttr) Marshal()([]byte, error){
	return json.Marshal(attr)
}

func (attr *FileAttr)UnMarshal(buf []byte) error{
	return json.Unmarshal(buf, attr)
}
