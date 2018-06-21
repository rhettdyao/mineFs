package nameserver

import (
	"testing"
	"github.com/syndtr/goleveldb/leveldb"
	"common"
)


var kdb *leveldb.DB
func init(){
	dbpath := "/tmp/leveldb"
	var err error
	kdb, err = leveldb.OpenFile(dbpath, nil)
	if err != nil{
		panic("open leveldb failed.")
	}
}

func TestDir(t *testing.T) {
	ns := NewNameSpace(kdb, kDefaultRootIno, 1)
	if ns == nil{
		t.Fatalf("failed to make namespace")
	}

	lists, code := ns.ReadDir("/")
	if code != common.EOK{
		t.Fatalf("can not read dir, ret %d", code)
	}
	if len(lists) != 0{
		t.Fatalf("the lists's size is abnormal, %v", len(lists))
	}

	code = ns.Mkdir("/home", 0755)
	if code != common.EOK{
		t.Fatalf("can not mkdir /home , ret=%d", code)
	}

	code = ns.Mkdir("/home/usr1", 755)
	if code != common.EOK{
		t.Fatalf("can not mkdir /home/usr1 , ret=%d", code)
	}
	code = ns.Mkdir("/home/usr2", 755)
	if code != common.EOK{
		t.Fatalf("can not mkdir /home/usr2 , ret=%d", code)
	}
	code = ns.Mkdir("/home/usr3", 755)
	if code != common.EOK{
		t.Fatalf("can not mkdir /home/usr3 , ret=%d", code)
	}
	list1, code := ns.ReadDir("/")
	if code != common.EOK{
		t.Fatalf("can not read dir, ret %d", code)
	}
	if len(list1) != 1 || list1[0] != "home"{
		t.Fatalf("the lists's size is abnormal, %v", len(list1))
	}

	list2, code := ns.ReadDir("/home")
	if code != common.EOK{
		t.Fatalf("can not read dir /home, ret %d", code)
	}
	if len(list2) != 3{
		t.Fatalf("the lists's size is abnormal, %v", len(list2))
	}
	for  _, v := range list2{
		if v != "usr1" &&
			v != "usr2" &&
			v != "usr3"{
			t.Fatalf("the lists's size is abnormal, %v", len(list2))
		}
	}
	code = ns.RmDir("/home/usr3")
	if code != common.EOK{
		t.Fatalf("can not rm dir /home/usr3, ret %d", code)
	}
	list3, code := ns.ReadDir("/home")
	if code != common.EOK{
		t.Fatalf("can not read dir /home, ret %d", code)
	}
	if len(list3) != 2{
		t.Fatalf("the lists's size is abnormal, %v", len(list3))
	}
}
