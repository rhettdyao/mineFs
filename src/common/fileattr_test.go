package common

import (
	"testing"
	"os"
	"time"
	"common/glog"
)

func TestDir(t *testing.T){
	attr := FileAttr{
		Name: "home",
		Mode: os.ModeDir | os.ModePerm,
		Blocks: 102023242393,
		Ctime: time.Now().Unix(),
		Mtime:time.Now().Unix(),
		Atime:time.Now().Unix(),
		Uid: 0,
		Gid: 0,
		StIno: 1000,
		ParentStIno: 1,
	}

	if !attr.IsDir(){
		t.Failed()
	}

	encode, err := attr.Marshal()
	if err != nil{
		t.Failed()
	}

	glog.InfoOut("the debug is abnormal\n")
	glog.InfoOut(string(encode))

	var new_attr FileAttr
	err = new_attr.UnMarshal(encode)
	if err != nil{
		t.Failed()
	}
	glog.InfoOut(new_attr.Name, new_attr.StIno, new_attr.IsDir())
}

func TestSplistPath(t *testing.T){
	fname := "/home/root/file"
	lists := SplistPath(fname)
	if len(lists) != 3{
		t.Failed()
	}
	if lists[0] != "home" || lists[1] != "root" || lists[2] != "file"{
		t.Failed()
	}

	

}
