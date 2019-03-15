package main

import (
	"github.com/jnst/go-world/logger"
	"os"
)

func main() {
	log := logger.CreateLogger()
	log.Info("succeeded to create logger")
	log.Debugf("pid: %d", os.Getpid())
	log.Errorw("failed to sample code", "lang", "ja")
}
