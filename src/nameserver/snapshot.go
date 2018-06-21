package nameserver

func Snapshot(){

}

/*
import (
	"sync/atomic"
	"strconv"
)
var kDefaultSnapShot = int64(0)
var kSnapshot = int64(1000)

func (m *NSManager)RecoverSnapshot(info *FileInfo) bool{
	if v, ok := info.xattrs["iSnapshotid"]; ok{
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil || id < 1000{
			return false
		}
		atomic.CompareAndSwapInt64(&kSnapshot, kSnapshot, id)
		ns := NewNameSpace(m.db, info.Ino)
		m.roots[id] = ns
		return true
	}
	return false
}
*/