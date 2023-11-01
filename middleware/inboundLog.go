package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/labstack/echo/v4"
)

// This function is used to set InboundLog
func (m *MiddlewareManager) InboundLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		startTime := time.Now()
		req := c.Request()
		res := c.Response()

		// Get request body
		reqBody := []byte{}
		if c.Request().Body != nil {
			reqBody, _ = ioutil.ReadAll(req.Body)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

		var mapReq map[string]interface{}
		var compactReq *bytes.Buffer

		if err := json.Unmarshal(reqBody, &mapReq); err != nil {
			if req.URL.Path != "/" {
				m.logger.Error("error: failed attempt to unmarshal request body from InboundLog")
			}
		}
		if err == nil {
			compactReq := new(bytes.Buffer)
			if err = json.Compact(compactReq, reqBody); err != nil {
				if req.URL.Path != "/" {
					m.logger.Error("error:", err)
				}
			}
		}

		if err = next(c); err != nil {
			c.Error(err)
		}

		latency := float64(time.Since(startTime).Nanoseconds()/1e4) / 100.0

		m.logger.Infow("INBOUND LOG",
			"qiscus_app_id", configs.LoadEnv("QISCUS_APP_ID"),
			"ip", req.RemoteAddr,
			"method", req.Method,
			"user_agent", req.UserAgent(),
			"path", req.URL.Path,
			"body", compactReq.String(),
			"status", res.Status,
			"latency", latency,
		)

		return
	}
}
