package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)
//SKAPAR ALLA EVENT HANDLERS

// Create event
func (app *application) createEvent(c *gin.Context) {
	var event database.Event

	if err := c.ShouldBindBodyWithJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := app.models.Events.Insert(&event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event zz"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

//get all event

func (app *application) getAllEvents(c *gin.Context){
		events, err := app.models.Events.GetAll()

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events (GetallEvents)"})
		}

		//allt ok 
		c.JSON(http.StatusOK, events)
}




// Get event
func (app *application) getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nope, Invalid event ID (getEvent)"})
	}

	event, err := app.models.Events.Get(id)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found 404"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed retrieve event"})
	}

	c.JSON(http.StatusOK, event)
}

//Update event

func (app *application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID (updateEvent)"})
		return
	}

		existingEvent, err := app.models.Events.Get(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
			return
		}

		if existingEvent == nil {
			c.JSON(http.StatusNotFound,gin.H{"error": "Event not found"})
		}

		updatedEvent := &database.Event{}
		if err := c.ShouldBindJSON(updatedEvent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
				return
		}


		updatedEvent.Id = id

		if err := app.models.Events.Update(updatedEvent); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
			return
		}

		c.JSON(http.StatusOK, updatedEvent)

}

func (app *application) deleteEvent (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid event ID"})
	}

	if err := app.models.Events.Delete(id); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete event"})
	}

	c.JSON(http.StatusNoContent, nil)
}

//IMPLEMENTERING AV ALLA DATABAS METODER SOM VI ANROPAR I VÅRA HANDLERS