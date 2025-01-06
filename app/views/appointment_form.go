package views

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sdassow/fyne-datepicker"
)

func NewAppointmentFormView(win fyne.Window) *fyne.Container {
	title := widget.NewLabelWithStyle("Agendar Nueva Cita", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Inputs del formulario
	datePicker := datepicker.NewDatePicker(time.Now(), time.Monday, func(t time.Time, b bool) {})
	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "Nombre del Cliente"

	confirmButton := widget.NewButton("Confirmar", func() {
		// Lógica para guardar la cita
	})

	backButton := widget.NewButton("Regresar", func() {
		win.SetContent(NewHomeView(win, nil))
	})

	return container.NewVBox(title, nameEntry, datePicker, confirmButton, backButton)
}
