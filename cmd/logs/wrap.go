package logs

import "github.com/yasseldg/go-simple/logs/sLog"

func Wrap() {

	sLog.Info("Wrap %l %s", "nueva linea")

	sLog.Info(sLog.Lines("Wrap %l %s %l y una mas %l   esta indentada"), "nueva linea")

	sLog.Debug(sLog.Lines("debug Wrap: %d %l con lineas"), 1)

	sLog.Warn("warn Wrap: %s", "atencion")

	sLog.Error("error Wrap (%d): %s", 2, "ninguno")

	sLog.Fatal("fatal Wrap: %s = %t", "grave", false)
}
