package slf4go_native_adaptor

import (
	"errors"
	"github.com/aellwein/slf4go"
	"testing"
)

func TestGetLogger(t *testing.T) {
	logger := slf4go.GetLogger("test")
	logger.SetLevel(slf4go.LevelTrace)
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

	// will actually exit and break the test
	//logger.Fatalf("import cycle not allowed! %s", "shit...")
	//logger.Fatal("never reach here")
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

func TestLoggerFormat(t *testing.T) {
	logger := slf4go.GetLogger("test")
	logger.Tracef("arr: %v, %d, %s", []int{1, 2, 3}, 102, "haha")
	logger.Tracef("arr: %d, %d, %f", 123, 102, 122.33)
}

/**
  BenchmarkLoggerCheckEnable-8      	500000000	         3.16 ns/op	       0 B/op	       0 allocs/op
  BenchmarkLoggerNotCheckEnable-8   	50000000	        32.9 ns/op	      16 B/op	       1 allocs/op
*/
func BenchmarkLoggerCheckEnable(b *testing.B) {
	logger := slf4go.GetLogger("test")
	logger.SetLevel(slf4go.LevelInfo)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if logger.IsTraceEnabled() {
			logger.Tracef("this is a test, b: %v, ", b)
		}
	}
}
func BenchmarkLoggerNotCheckEnable(b *testing.B) {
	logger := slf4go.GetLogger("test")
	logger.SetLevel(slf4go.LevelInfo)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		logger.Tracef("this is a test, b: %v, ", b)
	}
}

func TestSetLoggingParameters(t *testing.T) {
	if nil != slf4go.GetLoggerFactory().SetLoggingParameters(slf4go.LoggingParameters{}) {
		t.Error("expected no error")
	}
}
