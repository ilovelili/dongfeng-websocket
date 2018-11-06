package main

import (
	"fmt"
	"time"

	logger "github.com/ilovelili/dongfeng-logger"
	"github.com/ilovelili/dongfeng-websocket/services/server/app"
)

func main() {
	app := &app.App{}
	if err := app.Bootstarp(); err != nil {
		errorlog := &logger.Log{
			Category: "ErrorLog",
			Content:  fmt.Sprintf("WebSocket server bootstrap failed: %s\n", err.Error()),
			Time:     time.Now(),
		}

		errorlog.ErrorLog(logger.WebSocket)
	}
}
