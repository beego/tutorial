package main

import (
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {
	log := logs.NewLogger(10000)
	log.SetLogger("console", "")
	log.EnableFuncCallDepth(true) // 开启文件和行号显示
	log.Trace("trace %s %s", "param1", "param2")
	log.Debug("debug")
	log.Info("info")
	log.Warn("warning")
	log.Error("error")
	log.Critical("critical")

	time.Sleep(1 * 1e9) // why ?
}
