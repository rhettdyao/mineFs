package nameserver

import (
	"../github.com/rhettdyao/goleveldb/leveldb"
	"../github.com/rhettdyao/goleveldb/leveldb/util"
	"encoding/binary"
	"path/filepath"
	"os/exec"
	"sync/atomic"
	"container/list"
)

const kRooetEntryId = uint64(1)

var kCurretEntryId = kRooetEntryId

func encodingKeyStore(id uint64, name string) []byte{
	buf := make([]byte, 8 + len(name))
	binary.BigEndian.PutUint64(buf[:7], id)
	copy(buf[8:], []byte(name))
	return buf
}

func decodingKeyStore(k []byte)(uint64, string){
	id := binary.BigEndian.Uint64(k[:7])
	name := string(k[8:])
	return id, name
}

type NameSpace struct {
	db *leveldb.DB
	root_entry uint64
}

func (ns *NameSpace) getFromKeyStore(key []byte) (*FileInfo, bool){
	if value, err := ns.db.Get(key, nil); err != nil{
		var info FileInfo
		if info.DecodeFileInfo(value){
			return _, false
		}
		return &info, true
	}
	return _, false
}

func (ns *NameSpace)lookUpByParent(parent_id uint64, name string)(info *FileInfo, ret bool){
	return ns.getFromKeyStore(encodingKeyStore(parent_id, name))
}

func (ns *NameSpace) LookUp(name string)(info *FileInfo, ret bool){
	if name == "/"{
		return
	}

	split_list := filepath.SplitList(name)
	if len(split_list) == 0{
		return
	}

	parent_id, entry_id :=kRooetEntryId, kRooetEntryId
	for _, v := range split_list{
		if info, ret = ns.lookUpByParent(parent_id, v); ret != true{
			return
		}
		parent_id = entry_id
		entry_id = info.id
	}
	info.name = split_list[len(split_list) - 1]
	info.parent_id = parent_id
	return
}

func (ns *NameSpace)delFileInfo(key []byte) bool{
	return (ns.db.Delete(key, nil) == nil)
}

func (ns *NameSpace)UpdateFileInfo(info *FileInfo) bool{
	return (ns.db.Put(encodingKeyStore(info.parent_id, info.name), info.EncodeFileInfo(), nil) == nil)
}

func (ns *NameSpace)GetFileInfo(name string) (info *FileInfo, ret bool){
	if info, ret = ns.LookUp(name); ret == false{
		return
	}
	return
}

func (ns *NameSpace) buildPath(file_path string)(info *FileInfo, name string, ret bool){
	split_list := filepath.SplitList(name)
	parent_id := kRooetEntryId
	depth := len(file_path)

	for i := 0; i < len(split_list) -1;  i++{
		var file_info *FileInfo
		if file_info, ret = ns.lookUpByParent(parent_id, split_list[i]); ret == true {
			parent_id = file_info.id
		}else {
			return
		}
	}
	name = split_list[len(split_list) -1]
	return
}

func (ns *NameSpace) CreateFile(name string) (err error){
	if name == "/"{
		return
	}
	file_info, file_name, ret := ns.buildPath(name)

	if !ret{
		return
	}

	parent_id := file_info.parent_id
	_, exist := ns.lookUpByParent(parent_id, file_name)
	if (exist){
		return
	}

	var info FileInfo
	info.id = atomic.AddUint64(&kCurretEntryId, 1)
	info.parent_id = parent_id
	err = ns.db.Put(encodingKeyStore(info.parent_id, file_name), info.EncodeFileInfo(), nil)
	if err != nil{
		return
	}
	return
}

func (ns *NameSpace)ListDirectory(dir string) (childrens []string, err error){
	file_info, ret := ns.LookUp(dir)
	if ret == false{
		return
	}
	var r leveldb.util.Range
	r.Start = encodingKeyStore(file_info.id, "")
	r.Limit = encodingKeyStore(file_info.id+1, "")

	var temp list.List
	it  := ns.db.NewIterator(&r, nil)
	for it.Valid(){
		var info FileInfo
		ret := info.DecodeFileInfo(it.Value)
		if ret == false{
			return
		}
		temp.PushBack(info.name)
	}
	childrens = make([]string, temp.Len())
	var index = 0
	for el := temp.Front(); el != nil; el = el.Next(){
		childrens[index] = el.Value.(string)
		index++
	}

	return childrens, nil
}


func (ns *NameSpace) Rename(old_path, new_path string) (err error){
	if old_path == "/" || new_path == "/" || old_path == new_path{
		return nil
	}

	old_info, ret := ns.LookUp(old_path)
	if !ret{
		return nil
	}

	_, ret = ns.LookUp(new_path)
	if ret == true{
		return nil //false
	}
	dir, dst_name := filepath.Split(new_path)
	new_dir_info,ret := ns.LookUp(dir)
	if ret == false{
		return nil //false
	}
	old_info.parent_id = new_dir_info.id

	return ns.db.Put(encodingKeyStore(old_info.id, dst_name), old_info.EncodeFileInfo(), nil)
}

func (ns *NameSpace)RemoveFile(fname string) (err error){
	info, ret := ns.LookUp(fname)
	if !ret{
		return nil //success
	}
	
}



