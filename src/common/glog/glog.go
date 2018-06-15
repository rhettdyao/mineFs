package glog

import (
	"log"
	"os"
	"time"
	"sync"
	"fmt"
	"io"
)

const (
	LogLevelDebug = 1
	LogLevelInfo = 2
	LogLevelWarnning = 4
	LogLevelError = 8
)

type Logger struct {
	mu sync.Mutex
	loglevel int
	logdir string
	debug *log.Logger
	info *log.Logger
	warning *log.Logger
	error *log.Logger
	writer io.Writer
}

func init(){
	logger.init(LogLevelInfo | LogLevelWarnning | LogLevelError, "")
}

func logFileName() string{
	t := time.Now()
	temp := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	return "minelog_" + temp+".log"
}


func (logger *Logger)resetLogger() bool{
	logger.debug = nil
	logger.warning = nil
	logger.error = nil
	logger.info = nil

	var writer io.Writer = os.Stdout
	if logger.logdir != ""{
		var err error
		writer, err = os.OpenFile(logger.logdir + "/" + logFileName(), os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0755)
		if err != nil{
			return false
		}
		writer = io.MultiWriter(os.Stdout, writer)
	}


	logger.debug = log.New(writer, "[Debug] ", log.Ltime | log.Lshortfile)
	logger.info = log.New(writer, "[Info] ", log.Ltime | log.Lshortfile)
	logger.warning = log.New(writer, "[Warning] ", log.Ltime | log.Lshortfile)
	logger.error = log.New(writer, "[Error] ", log.Ltime | log.Lshortfile)
	logger.writer = writer
	return true
}

func (logger *Logger)init(level int, dir string) bool{
	logger.loglevel = level
	logger.logdir = dir
	return logger.resetLogger()
}

func (logger *Logger)LogOut(level int, v ...interface{}){
	logger.mu.Lock()
	defer logger.mu.Unlock()

	var log *log.Logger
	switch level {
	case level & logger.loglevel :
		log = logger.debug
	case level & LogLevelInfo:
		log = logger.info
	case level & LogLevelWarnning:
		log = logger.warning
	case level & LogLevelError:
		log = logger.error
	default:
		log = nil
	}
	msg := fmt.Sprint(v...)
	if log != nil{
		log.Output(3,msg)
	}
}

var logger Logger

type Verbose bool

func ErrorOut(v ...interface{}){
	logger.LogOut(LogLevelError, fmt.Sprint(v...))
}

func WarningOut(v ...interface{}){
	logger.LogOut(LogLevelWarnning, fmt.Sprint(v...))
}

func InfoOut(v ...interface{}){
	logger.LogOut(LogLevelInfo, fmt.Sprint(v...))
}

func DebugOut(v ...interface{}){
	logger.LogOut(LogLevelDebug, fmt.Sprint(v...))
}


func SetLog(level int, dir string) bool{
	return logger.init(level, dir)
}
