package nameserver

import "sync/atomic"

const kDefaultRootIno = int64(1)
const  kBeginIno = int64(1000)
var kCurrentIno = kBeginIno

func AutoGetIno() int64{
	return atomic.LoadInt64(&kCurrentIno)
}

func AutoAddIno() int64{
	return atomic.AddInt64(&kCurrentIno, 1)
}

func CompareAddBigger(ino int64){
	cur := AutoGetIno()
	if cur < ino{
		atomic.CompareAndSwapInt64(&kCurrentIno, cur, ino)
	}
}


