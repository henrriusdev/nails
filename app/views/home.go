package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/henrriusdev/nails/pkg/service"
)

func NewHomeView(win fyne.Window, customerService *service.Customer) *fyne.Container {
	title := widget.NewLabelWithStyle("Bienvenido al Spa de Uñas", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	agendarButton := widget.NewButtonWithIcon("Agendar Cita", theme.ContentAddIcon(), func() {
		win.SetContent(NewAppointmentFormView(win))
	})
	verCitasButton := widget.NewButtonWithIcon("Ver Citas", theme.AccountIcon(), func() {
		// win.SetContent(NewAppointmentsListView(win))
	})

	historialClientesButton := widget.NewButtonWithIcon("Clientes", theme.DocumentIcon(), func() {
		// win.SetContent(NewCustomerFormView(win, customerService))
	})

	return container.NewVBox(title, agendarButton, verCitasButton, historialClientesButton)
}
