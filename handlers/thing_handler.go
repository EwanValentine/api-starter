package handlers

import (
	"github.com/EwanValentine/api-starter/models"
	"github.com/labstack/echo"
	"net/http"
)

type Error struct {
	Message string `json:"_message"`
	Code    int    `json:"_code"`
}

type Response struct {
	Data interface{}            `json:"data"`
	Meta map[string]interface{} `json:"_meta"`
}

type ThingHandler struct {
	datastore *models.ThingRepository
}

func NewHandler(datastore *models.ThingRepository) *ThingHandler {
	return &ThingHandler{
		datastore,
	}
}

func (handler *ThingHandler) FindAll(c echo.Context) error {

	things, err := handler.datastore.FindAll()

	if err != nil {
		return c.JSON(404, &Error{
			Code:    http.StatusNotFound,
			Message: "No things found",
		})
	}

	return c.JSON(200, &Response{
		Data: things,
		Meta: map[string]interface{}{
			"_link": "/api/v1/things",
		},
	})
}
