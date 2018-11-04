package main

import (
	"time"

	logger "github.com/ilovelili/dongfeng-logger"
	"github.com/ilovelili/dongfeng-websocket/app"
)

func main() {
	app := &app.App{}
	if err := app.Bootstarp(); err != nil {
		syslog := &logger.SystemLog{
			Category:  "ErrorLog",
			Operation: "Web Socket Server Bootstrap",
			Content:   err.Error(),
			Time:      time.Now(),
		}

		util.ErrorLog(syslog, logger.Core_WebSocket)
	}
}
