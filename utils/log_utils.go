package utils

import (
	"os"
	"path"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"strings"
	"fmt"
	"runtime"
	"github.com/gin-gonic/gin"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
}

// config logrus log to local filesystem, with file rotation
func InitLogrus(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	if _, err := os.Stat(logPath); os.IsNotExist(err) { //如果不存在该目录，那么创建该目录
		os.MkdirAll(logPath, os.ModePerm)
	}
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, nil)

	logrus.AddHook(lfHook)
	logrus.AddHook(NewContextHook()) //添加行号

	logrus.SetFormatter(&logrus.TextFormatter{})
}


// ContextHook for log the call context
type contextHook struct {
	Field  string
	Skip   int
	levels []logrus.Level
}

// NewContextHook use to make an hook
func NewContextHook(levels ...logrus.Level) logrus.Hook {
	hook := contextHook{
		Field:  "line",
		Skip:   5,
		levels: levels,
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}
	return &hook
}

// Levels implement levels
func (hook contextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire implement fire
func (hook contextHook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.Field] = findCaller(hook.Skip)
	return nil
}

// 对caller进行递归查询, 直到找到非logrus包产生的第一个调用.
// 因为filename我获取到了上层目录名, 因此所有logrus包的调用的文件名都是 logrus/...
// 因此通过排除logrus开头的文件名, 就可以排除所有logrus包的自己的函数调用
func findCaller(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// 这里其实可以获取函数名称的: fnName := runtime.FuncForPC(pc).Name()
// 因为文件的全路径往往很长, 而文件名在多个包中往往有重复, 因此这里选择多取一层, 取到文件所在的上层目录那层.
func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	//fmt.Println(file)
	//fmt.Println(line)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}



func GinLogrusLogger() gin.HandlerFunc {
	skip := make(map[string]string)
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		ctx.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			// Stop timer
			end := time.Now()
			latency := end.Sub(start)
			clientIP := GetClientIP(ctx)
			method := ctx.Request.Method
			statusCode := ctx.Writer.Status()
			if raw != "" {
				path = path + "?" + raw
			}
			logrus.Infoln(fmt.Sprintf("[GIN] %v | %3d | %13v | %15s |%s %-7s",
				end.Format("2006/01/02 - 15:04:05"),
				statusCode,
				latency,
				clientIP,
				method,
				path,
			))
		}
	}
}

func GetClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		RemoteAddr := ctx.Request.Header.Get("Remote_addr")
		if RemoteAddr == "" {
			addr := strings.Split(ctx.Request.RemoteAddr, ":")
			return addr[0]
		} else {
			return RemoteAddr
		}
	}
	return ip
}