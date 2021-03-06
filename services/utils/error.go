package utils

import (
	"time"

	errorcode "github.com/ilovelili/dongfeng-error-code"
	logger "github.com/ilovelili/dongfeng-logger"
)

// NewError new rpc error
func NewError(customerror *errorcode.Error, detail ...string) error {
	e := customerror.NewError(detail...)

	go func() {
		errorlog := &logger.Log{
			Category: "ErrorLog",
			Content:  e.Error(),
			Time:     time.Now(),
		}
		errorlog.ErrorLog(logger.WebSocket)
	}()

	return e
}
