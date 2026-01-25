package endpoints

import (
	"encoding/json"
	"io"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

func NewLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()

		logLine := "REQUEST:"
		uri := req.RequestURI
		method := req.Method

		u, p, ok := req.BasicAuth()
		if ok {
			u = strings.TrimSpace(u)
			p = strings.TrimSpace(p)
			logLine += " Username=" + u + ", Password=" + p + ","
		}

		data, err := io.ReadAll(req.Body)
		if err != nil {
			c.Error(err)
			return next(c)
		}

		if len(data) == 0 {
			logLine += " Payload={}"
		} else {
			logLine += " Payload=" + string(data)
		}
		logLine += ", Method=" + method + ", URI=" + uri
		logLine += ", Headers="

		headerData, err := json.Marshal(req.Header)
		if err != nil {
			c.Error(err)
			return next(c)
		}

		logLine += string(headerData)
		log.Println(logLine)
		return next(c)
	}
}
