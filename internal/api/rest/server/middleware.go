package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func recoveryMiddleware(logger *zap.SugaredLogger) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(err)

					c.Error(echo.ErrInternalServerError)
				}
			}()
			return next(c)
		}
	}
}

func httpErrorHandler(logger *zap.SugaredLogger) func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		msg := http.StatusText(http.StatusInternalServerError)
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			msg = he.Error()
		}

		if code >= http.StatusInternalServerError {
			logger.Error(msg)
			msg = http.StatusText(http.StatusInternalServerError)
		}

		c.JSON(code, msg)
	}
}
