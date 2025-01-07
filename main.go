package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/henrriusdev/nails/app/views"
	"github.com/henrriusdev/nails/config"
	"github.com/henrriusdev/nails/pkg/repository"
	"github.com/henrriusdev/nails/pkg/service"
	"github.com/henrriusdev/nails/pkg/store"
)

func main() {
	a := app.NewWithID("com.example.nailsspa")
	w := a.NewWindow("Nails Spa")

	// Set window size for Android
	w.Resize(fyne.NewSize(360, 640))

	// initialize bd
	db, _ := store.NewConnection(config.Env)

	// Initialize the repos
	appointmentRepo := repository.NewAppointment(db)
	customerRepo := repository.NewCustomer(db)
	productRepo := repository.NewProduct(db)
	roleRepo := repository.NewRole(db)
	userRepo := repository.NewUser(db)

	services := service.Services{
		Appointment: service.NewAppointment(appointmentRepo),
		Customer:    service.NewCustomer(customerRepo),
		Product:     service.NewProduct(productRepo),
		Role:        service.NewRole(roleRepo),
		User:        service.NewUser(userRepo),
	}

	// Set the initial content to the home view
	w.SetContent(views.NewHomeView(w, services.Customer))
	w.ShowAndRun()
}
