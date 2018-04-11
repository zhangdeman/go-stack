package self_log

/**
 * 自定义log记录的配置文件
 */
const (
	debugLogFile  = "debug.log"  //debug日志
	noticeLogFile = "notice.log" //notice日志
	warnLogFile   = "warn.log"   //warning日志
	fatalLogFile  = "fatal.log"  //fatal日志
)

const (
	LogLevelDebug  = 4 //debug级别
	LogLevelNotice = 3 //notice级别
	LogLevelWarn   = 2 //warn级别
	LogLevelFatal  = 1 //fatal级别
)
