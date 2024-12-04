package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"ingrytech/internal/svcerr"

	"github.com/labstack/echo/v4"
)

func parsePositiveInt64(name, val string) (int64, error) {
	parsed, err := strconv.Atoi(val)
	if err != nil {
		return 0, svcerr.NewErr(svcerr.ErrBadRequest, fmt.Sprintf("invalid %s: %s", name, val))
	}

	if parsed <= 0 {
		return 0, svcerr.NewErr(svcerr.ErrBadRequest, fmt.Sprintf("%s must be positive: %s", name, val))
	}

	return int64(parsed), nil
}

func handleError(err error) error {
	var code int

	switch {
	case errors.Is(err, svcerr.ErrBadRequest):
		code = http.StatusBadRequest
	case errors.Is(err, svcerr.ErrNotFound):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}

	return echo.NewHTTPError(code, err.Error())
}
