package nameserver

import (
	"github.com/syndtr/goleveldb/leveldb"
	"common"
	"path/filepath"
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb/util"
	"sync/atomic"
	"fmt"
)

const (
	MODE_RDONLY = 0
	MODE_RDWR = 1
)

type NameSpace struct {
	root int64
	db *leveldb.DB
	mode int
	info FileInfo
}

func NewNameSpace(db *leveldb.DB, root int64, mode int) *NameSpace{
	var ns NameSpace
	ns.root = root
	ns.db = db
	ns.mode = mode
	ns.info.SetDefault("/", true)
	ns.info.Ino = root
	return &ns
}

func encodingKeyStore(id int64, name string) []byte{
	buf := make([]byte, 8 + len(name))
	binary.BigEndian.PutUint64(buf[:8], uint64(id))
	copy(buf[8:], []byte(name))
	return buf
}

func decodingKeyStore(k []byte)(int64, string){
	id := binary.BigEndian.Uint64(k[:8])
	name := string(k[8:])
	return int64(id), name
}


func (ns *NameSpace) putToDb(parent int64, name string, info *FileInfo) bool{
	key := encodingKeyStore(parent, name)
	data , err := info.Marshal()
	if err != nil{
		return false
	}
	err = ns.db.Put(key, data, nil)
	return (err == nil)
}

func (ns *NameSpace) delFromDb(parent int64, name string) bool{
	key := encodingKeyStore(parent, name)
	err := ns.db.Delete(key, nil)
	return err == nil
}

func (ns *NameSpace) getFromDb(parent int64, name string) (info *FileInfo, ret bool){
	info = &FileInfo{}
	key := encodingKeyStore(parent, name)
	data, err := ns.db.Get(key, nil)
	if err != nil{
		return nil, false
	}
	err = info.UnMarshal(data)
	if err != nil{
		return nil, false
	}
	return info, true
}

func (ns *NameSpace) lookupByParent(parent int64, name string) (*FileInfo, common.ErrorStatus){
	fmt.Printf("Trace: lookup file(%s) parent(%d) in\n", name, parent)
	v, err := ns.db.Get(encodingKeyStore(parent, name), nil)
	if err != nil{
		fmt.Printf("read file(%s) parent(%v) failed, err(%s)\n", name, parent, err.Error())
		return nil, common.ENotOk
	}

	var info FileInfo
	err = info.UnMarshal(v)
	if err != nil{
		fmt.Printf("unmarshal file failed, err(%s)\n", err.Error())
		return nil, common.ENotOk
	}
	return &info, common.EOK
}

func (ns *NameSpace) LookUp(fpath string) (info *FileInfo, errcode common.ErrorStatus){
	fmt.Printf("Trace: look up(%s) in\n", fpath)
	if fpath == "/"{
		return &ns.info, common.EOK
	}
	parent := ns.root
	lists := common.SplistPath(fpath)
	for i := 0; i < len(lists); i++{
		info, errcode = ns.lookupByParent(parent, lists[i])
		if errcode != common.EOK{
			return
		}
		parent = info.Ino
	}
	return info, common.EOK
}

func (ns *NameSpace) Mkdir(fpath string, perm int) common.ErrorStatus{
	fmt.Printf("Trace: Mk dir(%s) in\n", fpath)
	if fpath == "/" || fpath == ""{
		return common.EBadParam
	}
	if ns.mode == MODE_RDONLY{
		return common.ENoPermission
	}

	dir, name := filepath.Split(fpath)
	dir_info, errcode := ns.LookUp(dir)
	if errcode != common.EOK{
		return errcode
	}

	_, errcode = ns.lookupByParent(dir_info.Ino, name)
	if errcode == common.EOK{
		return common.EFileExist
	}

	var info FileInfo
	info.SetDefault(name, true)
	info.Ino = atomic.AddInt64(&kCurrentIno, 1)
	data, err := info.Marshal()
	if err != nil{
		fmt.Println("mashal failed")
		return common.ENotOk
	}

	err = ns.db.Put(encodingKeyStore(dir_info.Ino, name), data, nil)
	if err != nil{
		fmt.Println("put to db failed")
		return common.ENotOk
	}

	return common.EOK
}

func (ns *NameSpace) ReadDir(fpath string) (childrens []string, errcode common.ErrorStatus){
	fmt.Printf("Trace: read dir(%s) in\n", fpath)
	temp, errcode := ns.LookUp(fpath)
	if errcode != common.EOK{
		return
	}
	if !temp.IsDir(){
		return nil, common.ENotFound
	}
	r := util.BytesPrefix(encodingKeyStore(temp.Ino, ""))
	it := ns.db.NewIterator(r, nil)
	defer  it.Release()
	for it.First(); it.Valid(); it.Next(){
		var info FileInfo
		err := info.UnMarshal(it.Value())
		if err != nil{
			return nil, common.ENotOk
		}
		childrens = append(childrens, info.Name)
	}
	return childrens, common.EOK
}

func (ns *NameSpace) RmDir(fpath string) (errcode common.ErrorStatus){
	fmt.Printf("Trace: rm dir(%s) in\n", fpath)
	if fpath == "/"{
		return common.ENoPermission
	}

	childrens, errcode := ns.ReadDir(fpath)
	if errcode != common.EOK{
		return
	}
	if len(childrens) != 0{
		return common.EDirNotEmpty
	}

	info, errcode := ns.LookUp(fpath)
	if errcode != common.EOK{
		return
	}

	err := ns.db.Delete(encodingKeyStore(info.ParentIno, info.Name), nil)
	if err != nil{
		return common.ENotOk
	}
	return common.EOK
}

func (ns *NameSpace) Create(fpath string) (ret common.ErrorStatus){
	if fpath == "/"{
		return common.ENoPermission
	}

	dir, name := filepath.Split(fpath)
	dirinfo, ret := ns.LookUp(dir)
	if ret != common.EOK{
		return
	}
	_, ret = ns.lookupByParent(dirinfo.Ino, name)
	if ret == common.EOK{
		return common.EFileExist
	}
	var info FileInfo
	info.SetDefault(name, false)
	info.Ino = atomic.AddInt64(&kCurrentIno, 1)
	info.ParentIno = dirinfo.Ino

	if !ns.putToDb(info.ParentIno, info.Name, &info){
		return common.ENotOk
	}
	return common.EOK
}

func (ns *NameSpace) Delete(fpath string) (ret common.ErrorStatus){
	if fpath == "/"{
		return common.ENoPermission
	}
	info, ret := ns.LookUp(fpath)
	if ret != common.EOK{
		return
	}
	if !info.IsRegular(){
		return common.EFileExist
	}

	if !ns.delFromDb(info.ParentIno, info.Name){
		return common.ENotOk
	}
	return common.EOK
}


/*
type NSManager struct {
	lock sync.RWMutex
	roots map[int64] *NameSpace
	db *leveldb.DB
}

var Kmanager NSManager

func init(){
	dbpath := "/tmp"
	var err error
	Kmanager.db, err = leveldb.OpenFile(dbpath, nil)
	if err != nil{
		panic("Bug: can not open db\n")
	}

	Kmanager.roots[kDefaultSnapShot] = NewNameSpace(Kmanager.db, kDefaultRootIno)
	if !Kmanager.recover(){
		panic("Bug: Recover namespace failed\n")
	}
}

func (m *NSManager) recoverOne(info *FileInfo) bool{
	CompareAddBigger(info.Ino)
	if _, ok := info.xattrs["iroot"]; ok{
		if info.Ino == kDefaultRootIno{
			ns := NewNameSpace(m.db, info.Ino)
			m.roots[0] = ns
		}else {
			return m.RecoverSnapshot(info)
		}
	}
	return true
}

func (m *NSManager) recover() bool{
	it := m.db.NewIterator(nil, nil)
	for it.First(); it.Valid(); it.Next(){
		var info FileInfo
		if err := info.UnMarshal(it.Value()); err != nil{
			return false
		}
		if !m.recoverOne(&info){
			return false
		}
	}
	return true
}

func (m *NSManager) Get() bool{

}

*/


