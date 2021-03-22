package main

import (
	_ "github.com/dihedron/go-zap-utils/log"
	"go.uber.org/zap"
)

func main() {
	defer zap.L().Sync()
	zap.L().Debug("application starting...")
}
