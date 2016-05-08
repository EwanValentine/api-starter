package handlers

import (
	"github.com/EwanValentine/api-starter/models"
	"github.com/labstack/echo"
	"net/http"
)

type Null struct {
	Message string `json:"_message"`
}

type Response struct {
	Data interface{}            `json:"data"`
	Meta map[string]interface{} `json:"_meta"`
}

type Handler struct {
	datastore *models.Repository
}

func NewHandler(datastore *models.Repository) *Handler {
	return &Handler{
		datastore,
	}
}

func (handler *Handler) FindAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	things, err := handler.datastore.FindAll(limit, offset)

	if err != nil {
		return c.JSON(404, http.StatusNotFound)
	}

	return c.JSON(200, &Response{
		Data: things,
		Meta: map[string]interface{}{
			"_link": "/api/v1/things",
		},
	})
}
