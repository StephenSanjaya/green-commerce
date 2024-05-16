package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func MakeLogEntry(c echo.Context) *log.Entry {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
	})

	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
		"host":   c.Request().Host,
	})
}

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		MakeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = map[string]interface{}{
			"message":     report.Message,
			"status code": report.Code,
		}
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	MakeLogEntry(c).Error(report.Message)

	if report.Code == 14 {
		report.Code = 500
	}
	c.JSON(report.Code, report.Message.(map[string]interface{}))
}
