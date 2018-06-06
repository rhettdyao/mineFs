package main

import (
	"fmt"
	"time"
	"path"
	"encoding/binary"
	"os"
	"os/exec"
)

const rand_write_offset = uint64(20*1024*1024 + 597)
const rand_write_size = 64
const seq_write_size = 128*1024 - 5
var rand_buf = make([]byte, rand_write_size)
var seq_buf = make([]byte, seq_write_size)

func randFile(fp string) string{
	return path.Join(fp, fmt.Sprintf("randfile_", time.Now().Unix()))
}

func autoAddByte(buf []byte){
	temp := binary.BigEndian.Uint64(buf)
	temp++;
	binary.BigEndian.PutUint64(buf, temp)
}

func InitBuf(buf []byte) (err error){
	f, err := os.Open("/dev/srandom")
	if err != nil{
		return
	}
	defer f.Close()
	n1, err := f.Read(buf)
	if err != nil{
		return
	}
	fmt.Println("read buf size is: ", n1)
	return
}

func MakeBuf(size int)(buf []byte, err error){
	buf = make([]byte, size)
	InitBuf(buf)
	return
}

/*
 * 该函数的功能为随机写入一个数据然后顺序写入大量数据，为模拟vmdk的写入方式。
 */
func VmdkWriteTest(dir_name string, )(err error){
	file_name := randFile(dir_name)
	f, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		return
	}
	defer f.Close()
	if err = InitBuf(rand_buf); err != nil{
		return
	}
	if err = InitBuf(seq_buf); err != nil{
		return
	}
	max_write_count, write_seq_interval :=1024 * 1024, 10
	for i := 0; i < max_write_count; i++{
		_, err = f.Write(seq_buf)
		if err != nil{
			return
		}
		autoAddByte(seq_buf)
		if i % write_seq_interval == 0{
			autoAddByte(rand_buf)
			_, err = f.WriteAt(rand_buf, int64(rand_write_offset))
			if err != nil{
				return
			}

			f.Seek(0, os.SEEK_END)
		}
	}
	return
}

/*
 * 该函数的做法为随机多个点进行写入，然后顺序大量写入数据.方法和先前类似，但是有多个点的随机写入,考虑再加入一个读取吧。
 */
type RandomWriteInfo struct{
	off int64
	len int32
	buf []byte
}

/*
var rand_infos []RandomWriteInfo{
	{off: 1*1024+56, len: 23}

}
*/


/*
 *覆盖写入测试，反复写入一个文件
*/
func bfsRewrite(dir_name string, size int64, write_size int)(err error){
	buf := make([]byte, write_size)
	err = InitBuf(buf)
	if err != nil{
		return
	}
	file_name := randFile(dir_name)
	fd, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		return
	}
	defer fd.Close()

	stat := NewSpeedStat("rewrite", 10)
	for ; ;{
		fd.Seek(0, os.SEEK_SET)
		for i := int64(0); i < size; {
			var l int
			l, err = fd.Write(buf)
			if err != nil{
				return
			}
			autoAddByte(buf)
			stat.Upate(int64(l))
			i += int64(write_size)
		}
	}
	return
}

func bfsCreateFile(file_name string, file_size int64)(real_size int64, err error){
	const write_size = 128*1024
	buf := make([]byte, write_size)
	err = InitBuf(buf)
	if err != nil{
		return
	}
	fd, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		return
	}
	defer fd.Close()

	real_size = 0;
	for ; real_size < file_size; real_size += write_size{
		_, err = fd.Write(buf);
		if err != nil{
			return
		}
		autoAddByte(buf)
	}
	return
}

/*
 * 反复读取一个文件，不对文件进行校验，仅仅是反复读取文件， 不对文件进行校验。
 */
func bfsRead(dir_name string, file_size int64, read_size int)(err error){
	file_name := randFile(dir_name)
	file_size, err = bfsCreateFile(file_name, 40 * 1024*1024)
	if err != nil{
		return
	}
	fd, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		return
	}
	defer fd.Close()
	buf := make([]byte, read_size)
	InitBuf(buf)
	for {
		fd.Seek(0, os.SEEK_SET)
		for i := int64(0); i < file_size; i += int64(read_size) {
			_, err = fd.Read(buf)
			if err != nil {
				return
			}
		}
	}
}

func pathIsExist(fname string) bool{
	_, err := os.Stat(fname);
	if  err == nil || os.IsExist(err){
		return true
	}
	return false
}

/*
 * 大量文件测试，创建大量文件，并对文件写入少量数据。
 */
func bfsFileCreate(dir_name string, file_num int)(err error){
	dir_name = path.Join(dir_name, "file_create")
	if pathIsExist(dir_name){
		exec.Command("/bin/bash", "rm", dir_name, "-rf")
	}

	err = os.Mkdir(dir_name, 0755)
	if err != nil{
		return
	}
	for i := 0; i < file_num; i++{
		file_name := fmt.Sprintf("test_file_%d", i)
		content := fmt.Sprintf("this is file test. test id is %d", i)

		var fd *os.File
		fd, err = os.OpenFile(path.Join(dir_name, file_name), os.O_RDWR|os.O_CREATE, 0755)
		if err != nil{
			return
		}
		_, err = fd.Write([]byte(content))
		fd.Close()
	}
	return
}

/*
 * 设置多个随机点进行随机写入，同时开始顺序写入文件。每128K里面选择3个点来随机写入
 */
func VmdkWriteTest2(dir_name string, count int)(err error) {
	chunk_size := 1024 * 128
	offset_array := make([]int64, count*3)
	for i := 0; i < count; i++{
		offset_array[i * 3 + 0] = int64(i * chunk_size + 128)
		offset_array[i * 3 + 1] = int64(i * chunk_size + 5681)
		offset_array[i * 3 + 2] = int64(i * chunk_size + 13051)
	}

	buf, _ := MakeBuf(64)
	seq_buf, _ := MakeBuf(62*1024)

	file_name := randFile(dir_name)
	fd, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		return
	}

	st :=NewSpeedStat("VmdkWriteTest2", 10)
	for{
		fd.Seek(0, os.SEEK_END)
		for i := 0; i < 20; i++ {
			_, err = fd.Write(seq_buf)
			if err != nil {
				return
			}
			st.Upate(62*1024)
			autoAddByte(seq_buf)
		}

		for _, v := range offset_array{
			fd.WriteAt(buf, v)
			st.Upate(64)
		}
		autoAddByte(buf)
	}

}
