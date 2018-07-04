# Native adaptor for SLF4GO

This is a simplest adaptor implementation for 
[Simple Logging Facade for Go](https://github.com/aellwein/slf4go), it uses the native 
logging implementation from the "log" package.
 
An example usage is stupid simple:

```go

package main

import (
	"github.com/aellwein/slf4go"
	_ "github.com/aellwein/slf4go-native-adaptor"
)

func main() {
    logger := slf4go.GetLogger("mylogger")
    logger.Info("It works!")
    logger.Warnf("The answer is %d", 42)
}
```
Note the underscore in front of the import of the SLF4GO adaptor. 
 
You can change the logger implementation anytime, without changing the facade you
are using, only by changing the imported adaptor.