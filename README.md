# zap-log

A library to provide a quick and easy integration with Uber's Zap logger.

In order to integrate Uber's Zap logging library into your project, in your `main` package import this library anonynously, and then make sure that the logger is sync'ed when the application terminates, like this:

```golang

import (
    // ...
    _ "github.com/dihedron/go-zap-utils/log"
)

func main() {
    defer zap.L().Sync()

    // your code goes here
}
```