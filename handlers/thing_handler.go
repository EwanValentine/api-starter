package handlers

import (
	"github.com/EwanValentine/api-starter/models"
	"github.com/labstack/echo"
	"net/http"
)

// ThingHandler - accepts ThingRepository as arg
type ThingHandler struct {
	datastore *models.ThingRepository
}

// NewHandler - Creates new instance of ThingHandler
func NewThingHandler(datastore *models.ThingRepository) *ThingHandler {
	return &ThingHandler{
		datastore,
	}
}

// FindAll - Handler to find all the things
func (handler *ThingHandler) FindAll(c echo.Context) error {

	things, err := handler.datastore.FindAll()

	if err != nil {
		return c.JSON(http.StatusNotFound, NotFound)
	}

	return c.JSON(200, &Response{
		Data: things,
		Meta: map[string]interface{}{
			"_link": "/api/v1/things",
		},
	})
}

// Insert - Handler to insert a thing
func (handler *ThingHandler) Insert(c echo.Context) error {
	var thing models.Thing

	c.Bind(&thing)

	err := handler.datastore.Insert(thing)

	if err != nil {
		return c.JSON(422, Unprocessable)
	}

	return c.JSON(http.StatusCreated, nil)
}
