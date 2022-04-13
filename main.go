package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/storage"
)

type AppointmentsServer struct {
	store *storage.AppointmentsStore
}

func NewAppointmentsServer() *AppointmentsServer {
	store := storage.NewAppointmentsStore()
	return &AppointmentsServer{store}
}

func (as *AppointmentsServer) createAppointmentHandler(c *gin.Context) {
	a := storage.Appointment{}
	err := c.ShouldBindJSON(&a)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	id := as.store.Insert(a)
	fmt.Println(id)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (as *AppointmentsServer) getAppointmentHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	appointment, err := as.store.GetAppointmentByID(storage.IDType(id))
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (as *AppointmentsServer) deleteAppointmentHandler(c *gin.Context) {
	id, errParse := strconv.Atoi(c.Params.ByName("id"))
	if errParse != nil {
		c.String(http.StatusBadRequest, errParse.Error())
		return
	}

	err := as.store.Delete(storage.IDType(id))
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
}

func (as *AppointmentsServer) getAllAppointments(c *gin.Context) {
	appointments := as.store.GetAllAppointments()
	c.JSON(http.StatusOK, appointments)
}

func main() {
	// REST API SERVER
	// Requests
	// /appointments/ GET - get all appointment
	// /appointments/<id> GET - get <id> appointment
	// /appointments/<id> DELETE - delete <id> appointment
	// /appointments/ POST - creates a record and return id of record
	router := gin.Default()
	appointmentsServer := NewAppointmentsServer()

	router.POST("/appointments/", appointmentsServer.createAppointmentHandler)
	router.GET("/appointments/:id", appointmentsServer.getAppointmentHandler)
	router.DELETE("/appointments/:id", appointmentsServer.deleteAppointmentHandler)
	router.GET("/appointments/", appointmentsServer.getAllAppointments)

	router.Run("localhost:" + "8888")
}
