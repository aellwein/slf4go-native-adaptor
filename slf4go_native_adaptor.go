package slf4go_native_adaptor

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/aellwein/slf4go"
)

const (
	levelTrace = "TRACE"
	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
	levelFatal = "FATAL"
	levelPanic = "PANIC"
	callDepth  = 2
)

type loggerAdaptorNative struct {
	slf4go.LoggerAdaptor
	tf   string
	flag int
}

func (l *loggerAdaptorNative) Trace(args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelTrace {
		str := fmt.Sprint(args...)
		l.output(callDepth, levelTrace, str)
	}
}

func (l *loggerAdaptorNative) Tracef(format string, args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelTrace {
		str := fmt.Sprintf(format, args...)
		l.output(callDepth, levelTrace, str)
	}
}

func (l *loggerAdaptorNative) Debug(args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelDebug {
		str := fmt.Sprint(args...)
		l.output(callDepth, levelDebug, str)
	}
}

func (l *loggerAdaptorNative) Debugf(format string, args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelDebug {
		str := fmt.Sprintf(format, args...)
		l.output(callDepth, levelDebug, str)
	}
}

func (l *loggerAdaptorNative) Info(args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelInfo {
		str := fmt.Sprint(args...)
		l.output(callDepth, levelInfo, str)
	}
}

func (l *loggerAdaptorNative) Infof(format string, args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelInfo {
		str := fmt.Sprintf(format, args...)
		l.output(callDepth, levelInfo, str)
	}
}

func (l *loggerAdaptorNative) Warn(args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelWarn {
		str := fmt.Sprint(args...)
		l.output(callDepth, levelWarn, str)
	}
}

func (l *loggerAdaptorNative) Warnf(format string, args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelWarn {
		str := fmt.Sprintf(format, args...)
		l.output(callDepth, levelWarn, str)
	}
}

func (l *loggerAdaptorNative) Error(args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelError {
		str := fmt.Sprint(args...)
		l.output(callDepth, levelError, str)
	}
}

func (l *loggerAdaptorNative) Errorf(format string, args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelError {
		str := fmt.Sprintf(format, args...)
		l.output(callDepth, levelError, str)
	}
}

func (l *loggerAdaptorNative) Fatal(args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelFatal {
		str := fmt.Sprint(args...)
		l.output(callDepth, levelFatal, str)
		os.Exit(1)
	}
}

func (l *loggerAdaptorNative) Fatalf(format string, args ...interface{}) {
	if l.GetLevel() <= slf4go.LevelFatal {
		str := fmt.Sprintf(format, args...)
		l.output(callDepth, levelFatal, str)
		os.Exit(1)
	}
}

func (l *loggerAdaptorNative) Panic(args ...interface{}) {
	str := fmt.Sprint(args...)
	l.output(callDepth, levelPanic, str)
	panic(levelPanic)
}

func (l *loggerAdaptorNative) Panicf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	l.output(callDepth, levelPanic, str)
	panic(str)
}

func (l *loggerAdaptorNative) output(calldepth int, level, s string) error {
	var file string
	var line int
	var ts = time.Now().Format(l.tf)
	if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		lastIndex := strings.LastIndex(file, "/")
		if lastIndex > 0 {
			file = file[lastIndex+1:]
		}
	}
	result := fmt.Sprintf("%-29s [%-5s] %s:%d %s\n", ts, level, file, line, s)
	_, err := slf4go.Writer.Write([]byte(result))
	return err
}

// native logger factory implementation
type nativeLoggerFactory struct {
	level slf4go.LogLevel
}

func newNativeLoggerFactory() slf4go.LoggerFactory {
	factory := &nativeLoggerFactory{level: slf4go.LevelInfo}
	return factory
}

// it should be private
func newNativeLogger(name string, lvl slf4go.LogLevel) *loggerAdaptorNative {
	logger := new(loggerAdaptorNative)
	logger.SetName(name)
	logger.SetLevel(lvl)
	logger.tf = "2006-01-02 15:04:05.999"
	logger.flag = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	return logger
}

func (f *nativeLoggerFactory) GetLogger(name string) slf4go.Logger {
	return newNativeLogger(name, f.level)
}

func (*nativeLoggerFactory) SetLoggingParameters(params slf4go.LoggingParameters) error {
	// for the native adaptor, currently no parameters are supported.
	return nil
}

func (f *nativeLoggerFactory) SetDefaultLogLevel(lvl slf4go.LogLevel) {
	f.level = lvl
}

func (f *nativeLoggerFactory) GetDefaultLogLevel() slf4go.LogLevel {
	return f.level
}
