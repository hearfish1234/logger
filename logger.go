package logger

import (
	"os"
	"log"
)

//初始化
type Log struct {
	fp string
	level int
}

const(
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

//初始化日志文件
func SetLogger(filePath string, level int) *Log {

	logger := new(Log)
	logger.fp = filePath
	logger.level = level
	return logger
}

//设置日志级别
func (logger *Log) SetLog(msg string, level int) error {
	
	fp, err := os.OpenFile(logger.fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer fp.Close()

	if level >= logger.level {
	
		switch level {
		case 1:
			setDebug(fp, msg)
			break
		case 2:
			setInfo(fp, msg)
			break
		case 3:
			setWarn(fp, msg)
			break
		case 4:
			setError(fp, msg)
			break
		case 5:
			setFatal(fp, msg)
			break
		}
	}

	return nil
}

//写日志
func setDebug(fp *os.File, msg string) {

	debugLog := log.New(fp, "[DEBUG]", log.Ldate|log.Ltime|log.Llongfile)
	debugLog.Println(msg)
}

func setInfo(fp *os.File, msg string) {

	infoLog := log.New(fp, "[INFO]", log.Ldate|log.Ltime|log.Llongfile)
	infoLog.Println(msg)
}

func setWarn(fp *os.File, msg string) {

	warnLog := log.New(fp, "[WARNING]", log.Ldate|log.Ltime|log.Llongfile)
	warnLog.Println(msg)
}

func setError(fp *os.File, msg string) {

	errorLog := log.New(fp, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile)
	errorLog.Println(msg)
}

func setFatal(fp *os.File, msg string) {

	fatalLog := log.New(fp, "[FATAL]", log.Ldate|log.Ltime|log.Llongfile)
	fatalLog.Println(msg)
}
