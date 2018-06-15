package directory

import "sync/atomic"

const kRootIno = uint64(1)

var kCurretIno = kRootIno

func GetIno() uint64{
	return atomic.AddUint64(&kCurretIno, 1)
}

func SetIno(ino uint64) bool{
	return atomic.CompareAndSwapUint64(&kCurretIno, kCurretIno, ino)
}