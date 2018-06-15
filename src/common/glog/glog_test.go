package glog

import (
	"testing"
	"fmt"
)

func TestStdOut(t *testing.T){
	if SetLog(LogLevelDebug | LogLevelInfo | LogLevelError | LogLevelWarnning, "") == false{
		t.Failed()
	}
	fmt.Println("log in")
	DebugOut("this is a log info\n")
	InfoOut("this is a warnning log")
	WarningOut("this is a error log")
	ErrorOut("this is a debug log")
}

func TestFileOut(t *testing.T){
	if SetLog(LogLevelDebug | LogLevelInfo | LogLevelError | LogLevelWarnning, "/tmp") == false{
		t.Failed()
	}
	fmt.Println("log in")
	DebugOut("this is a log info\n")
	InfoOut("this is a warnning log")
	WarningOut("this is a error log")
	ErrorOut("this is a debug log")
}