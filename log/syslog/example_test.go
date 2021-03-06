package syslog_test

import (
	"fmt"

	gosyslog "log/syslog"

	"github.com/chenleji/kit/log"
	"github.com/chenleji/kit/log/level"
	"github.com/chenleji/kit/log/syslog"
)

func ExampleNewLogger_defaultPrioritySelector() {
	// Normal syslog writer
	w, err := gosyslog.New(gosyslog.LOG_INFO, "experiment")
	if err != nil {
		fmt.Println(err)
		return
	}

	// syslog logger with logfmt formatting
	logger := syslog.NewSyslogLogger(w, log.NewLogfmtLogger)
	logger.Log("msg", "info because of default")
	logger.Log(level.Key(), level.DebugValue(), "msg", "debug because of explicit level")
}
