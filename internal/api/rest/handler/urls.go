package handler

import "github.com/labstack/echo/v4"

func (h *Handler) URLs() map[string]map[string]echo.HandlerFunc {
	return map[string]map[string]echo.HandlerFunc{
		"/books": {
			"GET":  h.books,
			"POST": h.createBook,
		},
		"/books/:id": {
			"GET":    h.book,
			"DELETE": h.deleteBook,
			"PUT":    h.updateBook,
		},
	}
}
