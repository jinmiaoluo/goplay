package main

import (
	log "github.com/inconshreveable/log15"
)

func main() {
	logger := log.New()
	logger.SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StderrHandler))
	//const (
	//	LvlCrit Lvl = iota
	//	LvlError
	//	LvlWarn
	//	LvlInfo
	//	LvlDebug
	//)
	logger.Debug("debug msg")
	logger.Info("info msg")
	logger.Warn("warn msg")
	logger.Error("error msg")
	logger.Crit("crit msg")

}
