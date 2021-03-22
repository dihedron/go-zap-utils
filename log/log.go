package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	configuration zap.Config
	restore       func()
)

func init() {
	app := strings.Replace(filepath.Base(os.Args[0]), ".exe", "", 1)
	var (
		err    error
		logger *zap.Logger
	)
	if content, err := ioutil.ReadFile(app + ".json"); err == nil {
		if err := json.Unmarshal(content, &configuration); err == nil {
			logger, err = configuration.Build()
			if err != nil {
				panic(fmt.Sprintf("error initialising logger: %v", err))
			}
			restore = zap.ReplaceGlobals(logger)
			zap.L().Info("application starting with custom log configuration")
			return
		}
	}

	configuration = zap.NewProductionConfig()
	configuration.Encoding = "json" // or "console"
	configuration.OutputPaths = []string{fmt.Sprintf("%s-%d.log", app, os.Getpid())}
	configuration.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, err = configuration.Build()
	if err != nil {
		panic(fmt.Sprintf("error initialising logger: %v", err))
	}
	restore = zap.ReplaceGlobals(logger)
	zap.L().Info("application starting with default log configuration")
	return
}

func SetLevel(level zapcore.Level) {
	configuration.Level.SetLevel(level)

}
