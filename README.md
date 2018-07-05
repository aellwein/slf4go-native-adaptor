[![Go Report Card](https://goreportcard.com/badge/github.com/aellwein/slf4go-native-adaptor)](https://goreportcard.com/report/github.com/aellwein/slf4go-native-adaptor)
[![Coverage Status](https://img.shields.io/coveralls/github/aellwein/slf4go-native-adaptor/master.svg)](https://coveralls.io/github/aellwein/slf4go-native-adaptor?branch=master)
[![Build Status](https://img.shields.io/travis/aellwein/slf4go-native-adaptor/master.svg)](https://travis-ci.org/aellwein/slf4go-native-adaptor) 


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


## Logging parameters

None supported.