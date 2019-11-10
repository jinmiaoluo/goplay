package main

import (
	log "github.com/inconshreveable/log15"
)

func main() {
	// all loggers can have key/value context
	//srvlog := log.New("module", "app/server")

	// all log messages can have key/value context
	//m := make(map[string]string)
	//m["hello"] = "world"
	//srvlog.Warn("abnormal conn rate", "message", m["hello"])
	//srvlog.Crit("crit log")
	//srvlog.Debug("debug log")
	//srvlog.Error("error log")
	//srvlog.Info("info log")

	logger := log.New()
	logger.Warn("warn message")
	logger.Debug("debug message")
	logger.Error("error message")
	//
	//	// child loggers with inherited context
	//	connlog := srvlog.New("raddr", c.RemoteAddr())
	//	connlog.Info("connection open")
	//
	//	// lazy evaluation
	//	connlog.Debug("ping remote", "latency", log.Lazy{pingRemote})
	//
	//	// flexible configuration
	//	srvlog.SetHandler(log.MultiHandler(
	//		log.StreamHandler(os.Stderr, log.LogfmtFormat()),
	//		log.LvlFilterHandler(
	//			log.LvlError,
	//			log.Must.FileHandler("errors.json", log.JsonFormat()))))
}
