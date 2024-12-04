package handler

import (
	"net/http"
	"time"

	"ingrytech/internal/models"
	"ingrytech/internal/svcerr"

	"github.com/labstack/echo/v4"
)

type BookRequest struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationDate string `json:"publication_date"`
}

type BookResponse struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationDate string `json:"publication_date"`
}

// POST one
func (h *Handler) createBook(c echo.Context) error {
	var req BookRequest
	if err := c.Bind(&req); err != nil {
		return handleError(err)
	}

	book, err := convertBookToModel(req)
	if err != nil {
		return handleError(err)
	}

	res, err := h.app.CreateBook(c.Request().Context(), book)
	if err != nil {
		return handleError(err)
	}

	return c.JSON(http.StatusCreated, echo.Map{"book": res})
}

// GET all
func (h *Handler) books(c echo.Context) error {
	books, err := h.app.Books(c.Request().Context())
	if err != nil {
		return handleError(err)
	}

	res := make([]BookResponse, 0, len(books))
	for _, book := range books {
		res = append(res, convertBookToResponse(book))
	}

	return c.JSON(http.StatusOK, echo.Map{"books": res})
}

// GET one
func (h *Handler) book(c echo.Context) error {
	id, err := parsePositiveInt64("id", c.Param("id"))
	if err != nil {
		return handleError(err)
	}

	book, err := h.app.Book(c.Request().Context(), id)
	if err != nil {
		return handleError(err)
	}

	res := convertBookToResponse(book)

	return c.JSON(http.StatusOK, echo.Map{"book": res})
}

// PUT one
func (h *Handler) updateBook(c echo.Context) error {
	var req BookRequest
	if err := c.Bind(&req); err != nil {
		return handleError(err)
	}

	book, err := convertBookToModel(req)
	if err != nil {
		return handleError(err)
	}

	id, err := parsePositiveInt64("id", c.Param("id"))
	if err != nil {
		return handleError(err)
	}

	book.ID = id

	book, err = h.app.UpdateBook(c.Request().Context(), book)
	if err != nil {
		return handleError(err)
	}

	res := convertBookToResponse(book)

	return c.JSON(http.StatusOK, echo.Map{"book": res})
}

// DELETE one
func (h *Handler) deleteBook(c echo.Context) error {
	id, err := parsePositiveInt64("id", c.Param("id"))
	if err != nil {
		return handleError(err)
	}

	if err := h.app.DeleteBook(c.Request().Context(), id); err != nil {
		return handleError(err)
	}

	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}

func convertBookToModel(req BookRequest) (models.Book, error) {
	pubDate, err := time.Parse(time.RFC3339, req.PublicationDate)
	if err != nil {
		return models.Book{}, svcerr.NewErr(svcerr.ErrBadRequest, "invalid publication date")
	}

	return models.Book{
		Name:            req.Name,
		Author:          req.Author,
		PublicationDate: pubDate,
	}, nil
}

func convertBookToResponse(book models.Book) BookResponse {
	return BookResponse{
		ID:              book.ID,
		Name:            book.Name,
		Author:          book.Author,
		PublicationDate: book.PublicationDate.String(),
	}
}
