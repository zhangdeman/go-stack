package self_log

import (
	"os"
	"fmt"
	"strings"
	"time"
)

type Log struct {
	logPath string	//日志路径
	logLevel int	//日志记录级别
}

/**
 * 创建log实例
 * @param logPath 日志路径
 * @param level 日志级别
 */
func MakeLog(logPath string, level int) (Log, error) {
	log := Log{
		logPath:logPath,
		logLevel:level,
	}

	getPathInfo, _ := PathExists(logPath)
	if !getPathInfo {
		//目录不存在,创建
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			//目录创建失败
			fmt.Println("日志目录 " + logPath + " 创建失败")
			return log, err
		}
	}
	//创建日志文件
	logList := []string{
		logPath+debugLogFile,
		logPath+noticeLogFile,
		logPath+warnLogFile,
		logPath+fatalLogFile,
	}
	for _, fullFilePath := range logList{
		if !CheckFileIsExist(fullFilePath) {
			//文件不存在
			fmt.Println("日志文件不存在,开始创建创建 : "+fullFilePath)
			file, err := os.Create(fullFilePath)
			if err != nil {
				fmt.Println("日志文件创建失败 : "+fullFilePath)
			} else {
				fmt.Println("日志文件创建成功 : "+fullFilePath)
				file.Close()
			}
		}
	}

	return log, nil
}

/**
 * 判断文件夹是否存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 * 是否需要记录当前级别日志
 */
func (log *Log) allowLogLevel(level int) bool {
	if level <= log.logLevel {
		return true
	}
	return false
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
 * 向文件中写入日志
 */
func (log *Log) write(fullFilePath string, logStrList []string) (bool, error) {
	var err error
	var file *os.File
	file, err = os.OpenFile(fullFilePath, os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("日志写入失败 : "+fullFilePath)
		return false, err
	}

	//关闭文件
	defer file.Close()

	baseStr := "["+time.Now().String()+"]"
	logStr := baseStr + strings.Join(logStrList, "###") + "\n"
	file.WriteString(logStr)
	return true, nil
}

/**
 * 记录debug日志
 */
func (log *Log) WriteDebugStr(logStr ...string)  {
	if log.allowLogLevel(LogLevelDebug) {
		fullFilePath := log.logPath+debugLogFile
		log.write(fullFilePath, logStr)
	}
}

/**
 * 记录notice日志
 */
func (log *Log) WriteNoticeStr(logStr ...string)  {
	if log.allowLogLevel(LogLevelNotice) {
		fullFilePath := log.logPath+noticeLogFile
		log.write(fullFilePath, logStr)
	}
}

/**
 * 记录warn日志
 */
func (log *Log) WriteWarnStr(logStr ...string)  {
	if log.allowLogLevel(LogLevelWarn) {
		fullFilePath := log.logPath+warnLogFile
		log.write(fullFilePath, logStr)
	}
}

/**
 * 记录fatal日志
 */
func (log *Log) WriteFatalStr(logStr ...string)  {
	if log.allowLogLevel(LogLevelFatal) {
		fullFilePath := log.logPath+fatalLogFile
		log.write(fullFilePath, logStr)
	}
}