package slf4go_native_adaptor

import "github.com/aellwein/slf4go"

func init() {
	slf4go.SetLoggerFactory(newNativeLoggerFactory())
}
