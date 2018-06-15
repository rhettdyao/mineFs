package directory

import (
	"testing"
	"github.com/syndtr/goleveldb/leveldb"
)

var test_dbpath = "/tmp"
var ldb *leveldb.DB

func initTestSuit() bool{
	if ldb != nil{
		return true
	}
	var err error
	ldb, err = leveldb.OpenFile(test_dbpath, nil)
	if err != nil{
		return false
	}
	return true
}

func TestMakeRoot(t *testing.T){
	if !initTestSuit(){
		t.Failed()
	}

	root := NewDirectory(kRootIno, ldb)
}
