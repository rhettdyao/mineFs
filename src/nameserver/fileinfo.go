package nameserver

import (
	"os"
	"encoding/json"
)



type FileInfo struct {
	Name string
	Ino int64  "ino"
	ParentIno int64 "parent ino"
	Size int64 "the file size"
	Blocks int64 "the blocks"
	Mode os.FileMode  "the file mode"
	Uid int32
	Gid int32
	xattrs map[string]string
}

func (info *FileInfo) IsDir() bool{
	return info.Mode & os.ModeDir != 0
}

func (info *FileInfo) IsRegular() bool{
	return info.Mode & os.ModeDir == 0
}

func (info *FileInfo) Perm() os.FileMode{
	return info.Mode &os.ModePerm
}

func (info *FileInfo) SetPerm(perm os.FileMode){
	info.Mode = (info.Mode >> 16) | (perm <<16)
}

func (info *FileInfo) SetDir(){
	info.Mode = info.Mode | os.ModeDir
}

func (info *FileInfo) SetRegular(){
	info.Mode = info.Mode & os.ModePerm
}

func (info *FileInfo) SetDefault(name string, isdir bool){
	info.Name = name
	info.Size = 0
	if isdir{
		info.Size = 4096
		info.Mode = os.ModeDir | 0755
	}else {
		info.Mode = 0755
	}
}

func (info *FileInfo) Marshal()([]byte, error){
	return json.Marshal(info)
}

func (info *FileInfo)UnMarshal(buf []byte) error{
	return json.Unmarshal(buf, info)
}


