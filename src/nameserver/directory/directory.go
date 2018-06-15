package directory

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"encoding/binary"
	"common"
	"time"
	"os"
	"path/filepath"
)

var db_path = ""

func init(){
	db_path = "/tmp"
}

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

type Directory struct {
	root uint64
	db *leveldb.DB
	attr common.FileAttr
}

func NewDirectory(root uint64, db *leveldb.DB) *Directory{
	var dir Directory
	dir.db = db
	dir.root = root
	dir.attr.Name = "/"
	dir.attr.Mode = os.ModeDir | 0755

	return &dir
}

func (dir *Directory) getKeyFromStore(key []byte)(*common.FileAttr, bool){
	val, err := dir.db.Get(key, nil)
	if err != nil{
		return _, false
	}
	var attr common.FileAttr
	err = attr.UnMarshal(val)
	if err != nil{
		return _, false
	}
	return &attr, true
}

func (dir *Directory) deleteFileInfo(key []byte) bool{
	return dir.db.Delete(key, nil) == nil

}

func (dir *Directory) lookUpByParent(parent_id uint64, name string)(*common.FileAttr, bool){
	key := encodingKeyStore(parent_id, name)
	return dir.getKeyFromStore(key)
}

func (dir *Directory) LookUp(name string)(attr *common.FileAttr, ret bool){
	if name == "/"{
		temp := *attr
		return &temp, true
	}

	lists := common.SplistPath(name)
	if len(lists) == 0{
		return nil, false
	}
	parent_id, entry_id :=dir.root, dir.root
	for _, v := range lists {
		if attr, ret = dir.lookUpByParent(parent_id, v); ret != true{
			return
		}
		parent_id = entry_id
		entry_id = attr.StIno
	}

	attr.ParentStIno = parent_id
	return attr, true
}

func (dir *Directory)delFileInfo(key []byte) bool{
	return (dir.db.Delete(key, nil) == nil)
}

func (dir *Directory)UpdateFileInfo(attr *common.FileAttr) bool{
	data, ret := attr.Marshal()
	if ret != nil{
		return false
	}
	return (dir.db.Put(encodingKeyStore(attr.ParentStIno, attr.Name), data, nil) == nil)
}

func (dir *Directory)GetFileInfo(name string) (attr *common.FileAttr, ret bool){
	return dir.LookUp(name)
}

func (dir *Directory) buildPath(file_path string) (attr *common.FileAttr, name string, err common.ErrorStatus){
	lists := common.SplistPath(file_path)
	parent_id := dir.root
	depth := len(lists)

	for i = 0; i < depth -1; i++{
		var ret bool
		if attr, ret = dir.lookUpByParent(parent_id, lists[i]); ret{
			parent_id = attr.StIno
		}else {
			return attr, name, common.ETargetDirExist
		}
	}
	name = lists[len(lists) -1]
	return attr, name, common.EOK
}

func (dir *Directory) CreateFile(file_path string) (errcode common.ErrorStatus){
	if file_path == ""{
		return common.EBadParam
	}

	attr, name, errcode := dir.buildPath(file_path)
	if errcode != common.EOK{
		return errcode
	}
	parent_id := attr.ParentStIno
	_, ret := dir.lookUpByParent(parent_id, name)
	if ret{
		return common.EFileExist
	}

	var file_attr common.FileAttr
	file_attr.ParentStIno = parent_id
	file_attr.StIno = GetIno()
	file_attr.Name = name
	file_attr.Ctime = time.Now().Unix()
	file_attr.Mtime = time.Now().Unix()
	file_attr.Atime = time.Now().Unix()
	file_attr.Mode = 0755
	data, err := file_attr.Marshal()
	if err != nil{
		return common.ENotOk
	}
	err = dir.db.Put(encodingKeyStore(file_attr.ParentStIno, name), data, nil)
	if err != nil{
		return common.ENotOk
	}
	return common.EOK
}

func (dir *Directory)ListDirectory(file_path string) ([]string, common.ErrorStatus){
	attr, ret := dir.LookUp(file_path)
	if !ret{
		return _, common.ETargetDirExist
	}
	var childrens []string
	start := encodingKeyStore(attr.StIno, "")
	it  := dir.db.NewIterator(util.BytesPrefix(start), nil)
	for it.Valid(){
		var file_attr common.FileAttr
		file_attr.UnMarshal(it.Value())
		childrens = append(childrens, file_attr.Name)
	}
	return childrens, common.EOK
}

func (dir *Directory) Mkdir(file string)(common.ErrorStatus){
	
}


func (dir *Directory) Rename(old_path, new_path string)(common.ErrorStatus){
	if old_path == "/" || new_path == "/" || old_path == new_path{
		return common.EBadParam
	}

	old_attr, ret := dir.LookUp(old_path)
	if !ret{
		return common.EBadParam
	}

	_, ret = dir.LookUp(new_path)
	if ret{
		return common.EFileExist
	}

	dirname, fname := filepath.Split(new_path)
	new_dir_attr , ret := dir.LookUp(dirname)
	if !ret{
		return common.EBadParam
	}
	old_attr.ParentStIno = new_dir_attr.StIno
	old_attr.Name = fname
	data, err := old_attr.Marshal()
	if err != nil{
		return common.EBadParam
	}

	err = dir.db.Put(encodingKeyStore(old_attr.StIno, fname), data, nil)
	if err != nil{
		return common.EBadParam
	}
	return common.EOK
}

func (dir *Directory) RemoveFile(fname string) common.ErrorStatus{
	attr, ret := dir.LookUp(fname)
	if !ret{
		return common.ENotFound
	}
	if fname == "/"{
		return common.ENoPermission
	}

	if attr.IsDir(){
		return common.ENoPermission
	}

	data := encodingKeyStore(attr.ParentStIno, attr.Name)
	if dir.deleteFileInfo(data) != true{
		return common.ENotOk
	}

	return common.EOK
}


func (dir *Directory) internalRmDir(name string, id, parent_id uint64, recursive bool) common.ErrorStatus{
	it := dir.db.NewIterator(util.BytesPrefix(encodingKeyStore(id, "")), nil)
	if it.Valid(){
		if !recursive{
			return common.EDirNotEmpty
		}
		for ;it.Valid(); it.Next(){
			var attr common.FileAttr
			if attr.UnMarshal(it.Value()) != nil{
				return common.ENotOk
			}
			if attr.IsDir(){
				dir.internalRmDir(attr.Name, attr.StIno, attr.ParentStIno, true)
				continue
			}
			ret := dir.deleteFileInfo(encodingKeyStore(attr.ParentStIno, attr.Name))
			if !ret {
				return common.ENotOk
			}
		}
	}

	ret := dir.deleteFileInfo(encodingKeyStore(parent_id, name))
	if !ret{
		return common.ENotOk
	}
	return common.EOK
}

func (dir *Directory) Rmdir(name string) common.ErrorStatus{
	attr, ret := dir.LookUp(name)
	if !ret{
		return common.ENotFound
	}
	if !attr.IsDir(){
		return common.EBadParam
	}
	return dir.internalRmDir(attr.Name, attr.ParentStIno, attr.StIno, false);
}


