package mainApp

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
)

func InfoModict(a fyne.App, UserID string) {
	w := a.NewWindow("정보수정")
	w.Resize(fyne.NewSize(300, 200))

	title_L := canvas.NewText("정보수정", color.White)
	title_L.TextSize = 30

	var data []Data
	var TempData []Data
	User := fmt.Sprintf("User\\%s.txt", UserID)
	UserFile, _ := os.ReadFile(User)
	json.Unmarshal(UserFile, &data)

	UserPassword := data[0].Password
	UserName := data[0].Name
	UserPhone := data[0].Phone

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

		data = TempData

		result, _ := json.MarshalIndent(data, "", " ")

		os.WriteFile(User, result, 0644)

		completeWindow := fyne.CurrentApp().NewWindow("완료")
		CompleteLabel := widget.NewLabel("변경되었습니다.")
		CwBtn := widget.NewButton("확인", func() {
			w.Close()
			completeWindow.Close()
		})
		completeWindow.SetContent(container.NewVBox(CompleteLabel, CwBtn))
		completeWindow.Show()

	})

	Cancle := widget.NewButton("취소", func() {
		w.Close()
	})

	buttons := container.NewHBox(Complete, Cancle)

	w.SetContent(container.NewVBox(IdEntry, PassEntry, NameEntry, PhoneEntry, buttons))
	w.Show()
}
