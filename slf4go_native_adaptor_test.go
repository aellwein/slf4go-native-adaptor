package slf4go_native_adaptor

import (
	"errors"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/aellwein/slf4go"
	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	logger := slf4go.GetLogger("test")

	var levels = []slf4go.LogLevel{
		slf4go.LevelPanic,
		slf4go.LevelFatal,
		slf4go.LevelError,
		slf4go.LevelWarn,
		slf4go.LevelInfo,
		slf4go.LevelDebug,
		slf4go.LevelTrace,
	}

	for _, i := range levels {
		logger.SetLevel(i)
		logger.Trace("Trace")
		logger.Tracef("Tracef: %v", logger)
		logger.Debug("Debug")
		logger.Debugf("Debugf: %s", "debug mode")
		logger.Info("Info")
		logger.Infof("Infof: %v", slf4go.GetLoggerFactory())
		logger.Warn("Warn")
		logger.Warnf("Warnf: %d", 42)
		logger.Error("Error")
		logger.Errorf("Errorf: %v", errors.New("some error"))
	}
}

func TestLoggerFatal(t *testing.T) {
	mockExit := func(int) {
		panic("mockExit called")
	}
	patch := monkey.Patch(os.Exit, mockExit)
	defer patch.Unpatch()

	logger := slf4go.GetLogger("test")
	underTest := func() {
		logger.Fatal("fatality!")
	}

	assert.Panics(t, underTest, "should panic")
}

func TestLoggerFatalf(t *testing.T) {
	mockExit := func(int) {
		panic("mockExit called")
	}
	patch := monkey.Patch(os.Exit, mockExit)
	defer patch.Unpatch()

	logger := slf4go.GetLogger("test")
	underTest := func() {
		logger.Fatalf("fatality: %d", 42)
	}
	assert.Panics(t, underTest, "should panic")
}

func TestLoggerPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()
	logger := slf4go.GetLogger("test")
	logger.Panic("this is expected to cause panic")
}

func TestLoggerPanicf(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic as expected")
		}
	}()
	logger := slf4go.GetLogger("test")
	logger.Panicf("this is expected to cause panic: %d", 42)
}

func TestSetLoggingParameters(t *testing.T) {
	assert.Nil(t, slf4go.GetLoggerFactory().SetLoggingParameters(slf4go.LoggingParameters{}))
}

func TestNativeLoggerFactory_GetDefaultLogLevel(t *testing.T) {
	assert.Equal(t, slf4go.GetLoggerFactory().GetDefaultLogLevel(), slf4go.LevelInfo)
}

func TestNativeLoggerFactory_SetDefaultLogLevel(t *testing.T) {
	slf4go.GetLoggerFactory().SetDefaultLogLevel(slf4go.LevelTrace)
	logger := slf4go.GetLogger("test")
	assert.True(t, logger.IsTraceEnabled())
}
