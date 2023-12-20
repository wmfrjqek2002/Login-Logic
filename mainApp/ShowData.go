package mainApp

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"strings"
)

func showData(a fyne.App, fileName string, jsonData []Data) {
	filemultxt := strings.TrimSuffix(fileName, ".txt")
	window := a.NewWindow(filemultxt)

	var TempData []Data

	UserID := jsonData[0].ID
	UserPassword := jsonData[0].Password
	UserName := jsonData[0].Name
	UserPhone := jsonData[0].Phone

	IdEntry := widget.NewEntry()
	IdEntry.Text = UserID
	IdEntry.SetPlaceHolder(UserID)

	PassEntry := widget.NewEntry()
	PassEntry.Text = UserPassword
	PassEntry.SetPlaceHolder(UserPassword)

	NameEntry := widget.NewEntry()
	NameEntry.Text = UserName
	NameEntry.SetPlaceHolder(UserName)

	PhoneEntry := widget.NewEntry()
	PhoneEntry.Text = UserPhone
	PhoneEntry.SetPlaceHolder(UserPhone)

	Complete := widget.NewButton("확인", func() {
		NewData := &Data{
			ID:       IdEntry.Text,
			Password: PassEntry.Text,
			Name:     NameEntry.Text,
			Phone:    PhoneEntry.Text,
		}
		TempData = append(TempData, *NewData)

		jsonData = TempData

		result, _ := json.MarshalIndent(jsonData, "", " ")

		User := fmt.Sprintf("User\\%s", filemultxt)

		os.WriteFile(User, result, 0644)

		completeWindow := fyne.CurrentApp().NewWindow("완료")
		CompleteLabel := widget.NewLabel("변경되었습니다.")
		CwBtn := widget.NewButton("확인", func() {
			window.Close()
			completeWindow.Close()
		})
		completeWindow.SetContent(container.NewVBox(CompleteLabel, CwBtn))
		completeWindow.Show()

		window.Show()

	})

	Cancle := widget.NewButton("취소", func() {
		window.Close()
	})

	buttons := container.NewHBox(Complete, Cancle)

	window.SetContent(container.NewVBox(IdEntry, PassEntry, NameEntry, PhoneEntry, buttons))
	window.Show()

}
