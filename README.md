# zoggly (ARCHIVED)
Loggly hooks for GO uber :zap: Zap logger

## Archived
This is probably so out of sync with the main package there's no reason to have it now.

## Usage
This example sends to loggly and standard out.

```go
package main

import (
    "github.com/peppage/zoggly"
    "github.com/uber-go/zap"
)

var logglyToken string = "token from source setup"
var logger zap.Logger

func main() {
    zl := NewZapLoggly(logglyToken)
	logger = zap.New(zap.NewJSONEncoder(), zap.ErrorLevel,
		zap.Output(zap.Tee(os.Stdout, zl.GetWriter())))

    // Make sure to flush to send all logs before exiting
    defer zl.Flush()
}

```
