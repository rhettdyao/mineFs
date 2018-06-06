package main

import (
	"time"
	"fmt"
)

type SpeedStat struct{
	label string
	total_size int64
	cur_size int64
	total_io int64
	cur_io int64
	begin_time time.Time
	last_time time.Time
	show_timval int
}

func NewSpeedStat(label string, timeval int) SpeedStat{
	return SpeedStat{label:label, total_size:0, cur_size:0, begin_time:time.Now(), last_time:time.Now(), show_timval:timeval}
}


func stringSizeFloat64(size float64) string{
	var mb, kb float64 = 1024*1024, 1024
	if size/mb > 1{
		return fmt.Sprintf("%.02fMB", size/mb)
	}
	if size / kb > 1{
		return fmt.Sprintf("%.02fKB", size/kb)
	}
	return fmt.Sprintf("%.02fByte", size)
}

func stringSizeInt64(size int64) string{
	return stringSizeFloat64(float64(size))
}

func clacSpeed(size int64, begin_time time.Time) string{
	now := time.Now()
	speed := float64(size) / float64(now.Sub(begin_time).Seconds())
	if speed < 0{
		fmt.Println("invalid speed.")
	}
	return (stringSizeFloat64(speed) + "/s")
}

func clacIops(io int64, begin_time time.Time) string{
	now := time.Now()
	speed := float64(io) / float64(now.Sub(begin_time).Seconds())
	return fmt.Sprintf("%.02fiops/s", speed)
}

func (s SpeedStat) Show(){
	fmt.Printf("==================================show %s info=======================================\n", s.label)
	fmt.Printf("==total_size===time(sec)===speed===cur_size===cur_time(sec)===speed===t_iops===c_iops\n")
	fmt.Printf("%s %.02f %s %s %.02f %s %s %s\n", stringSizeInt64(s.total_size), time.Now().Sub(s.begin_time).Seconds(),clacSpeed(s.total_size, s.begin_time),
		stringSizeInt64(s.cur_size), time.Now().Sub(s.last_time).Seconds(),clacSpeed(s.cur_size, s.last_time), clacIops(s.total_io, s.begin_time),
		clacIops(s.cur_io, s.last_time))
}

func (s *SpeedStat)Upate(size int64){
	s.total_size += size
	s.cur_size += size
	s.cur_io++
	s.total_io++
	now := time.Now()
	if int(now.Sub(s.last_time).Seconds()) < s.show_timval{
		return
	}
	s.Show()
	s.last_time = now
	s.cur_size = 0
	s.cur_io = 0
}


func tt(){
	s := NewSpeedStat("test", 10)
	for  i := 0; i < 10000; i++{
		time.Sleep(10 * time.Millisecond)
		s.Upate(128*1024)
	}
}
