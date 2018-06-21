package nameserver

import (
	"testing"
)

func TestMode(t *testing.T){
	var info FileInfo
	info.Name = "home"
	info.SetDir()
	if !info.IsDir(){
		t.Failed()
	}
	if info.IsRegular(){
		t.Failed()
	}
	info.SetRegular()
	if info.IsDir(){
		t.Failed()
	}

	if !info.IsRegular(){
		t.Failed()
	}

	info.SetPerm(0755)
	if !info.IsRegular(){
		t.Failed()
	}
	if info.Perm() != 0755{
		t.Failed()
	}

	info.SetRegular()
	if !info.IsRegular(){
		t.Failed()
	}
	if info.Perm() != 0755{
		t.Failed()
	}

	info.SetDir()
	if !info.IsDir(){
		t.Failed()
	}
	if info.Perm() != 0755{
		t.Failed()
	}
}

func TestMashal(t *testing.T){
	var info FileInfo
	info.Name = "home"
	info.Ino = AutoAddIno()
	info.ParentIno = 1
	info.SetDefault("home", true)
	info.Size = 4096
	info.Blocks = 1024

	buf,err := info.Marshal()
	if err != nil{
		t.Failed()
	}
	var new_info FileInfo
	err = new_info.UnMarshal(buf)
	if err != nil{
		t.Failed()
	}

}
