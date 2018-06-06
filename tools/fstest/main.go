package main

import (
	"fmt"
	"os"
	"flag"
	"errors"
)

func help(){
	msg := "usage:\n" +
		   "appname --dir= --suit=\n" +
			"0 - unkown; 1-bfsRewrite; 2-bfsFileCreate; 3-VmdkWriteTest2; 4 - VmdkWriteTest"
	fmt.Println(msg)
	os.Exit(0)
}

func main(){
	suit_flag := flag.Int("suit", 0, "0 - unkown; 1-bfsRewrite; 2-bfsFileCreate; 3-VmdkWriteTest2; 4 - VmdkWriteTest")
	dir := flag.String("dir", "", "the test dir")
	if *dir == ""{
		help()
	}

	fmt.Println("the dir is:", *dir)

	var err error = nil
	switch *suit_flag {
	case 4:
		err = VmdkWriteTest("/mnt/test_bfs")
	case 3:
		err = VmdkWriteTest2("/mnt/test_bfs", 50)
	case 2:
		err = bfsFileCreate("/mnt/test_bfs", 100*1024)
	case 1:
		err = bfsRewrite("/mnt/test_bfs", 20*1024*1024*1024, 61*1024 + 79)
	default:
		err = errors.New("not support.")
	}

	if err != nil{
		fmt.Println("test failed: ", err.Error())
	}else {
		fmt.Println("test success.")
	}
	return
}
