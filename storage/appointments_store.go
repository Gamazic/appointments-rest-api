package storage

import "fmt"

type IDType int

type Appointment struct {
	Name    string `json:"name"`
	Disease string `json:"disease"`
}

type AppointmentsStore struct {
	store       *DB[Appointment]
	availableID IDType
}

func NewAppointmentsStore() *AppointmentsStore {
	store := &DB[Appointment]{}
	store.AssignEntries()
	return &AppointmentsStore{store, 0}
}

func (as *AppointmentsStore) setNewAvailableID() {
	as.availableID += 1
}

func (as *AppointmentsStore) hasID(id IDType) bool {
	return as.store.HasKey(id)
}

func (as *AppointmentsStore) Insert(a Appointment) IDType {
	insertedID := as.availableID
	as.store.Insert(insertedID, a)
	as.setNewAvailableID()
	return insertedID
}

func (as *AppointmentsStore) Delete(id IDType) error {
	err := as.store.Delete(id)
	if err != nil {
		return fmt.Errorf("appointment with id=%d not found", id)
	} else {
		return nil
	}
}

func (as *AppointmentsStore) GetAppointmentByID(id IDType) (Appointment, error) {
	appointment, err := as.store.Find(id)
	if err != nil {
		return appointment, fmt.Errorf("appointment with id=%d not found", id)
	} else {
		return appointment, err
	}
}

func (as *AppointmentsStore) GetAllAppointments() []Appointment {
	return as.store.Values()
}
