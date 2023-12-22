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
	window.Resize(fyne.NewSize(300, 300))

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

	Dlt := widget.NewButton("삭제", func() {
		confirmWindow := fyne.CurrentApp().NewWindow("확인")
		ConfirmLabel := widget.NewLabel("정말로 삭제하시겠습니까?")
		ConfirmYesBtn := widget.NewButton("예", func() {
			UserFilePath := fmt.Sprintf("User\\%s.txt", filemultxt)
			err := os.Remove(UserFilePath)
			if err != nil {
				fmt.Println("파일 삭제 오류:", err)
			} else {
				confirmWindow.Close()
				window.Close()
				// 여기에 필요한 추가 작업을 수행하세요.
			}
		})
		ConfirmNoBtn := widget.NewButton("아니오", func() {
			confirmWindow.Close()
		})
		confirmWindow.SetContent(container.NewVBox(ConfirmLabel, ConfirmYesBtn, ConfirmNoBtn))
		confirmWindow.Show()
	})

	buttons := container.NewHBox(Complete, Cancle, Dlt)

	window.SetContent(container.NewVBox(IdEntry, PassEntry, NameEntry, PhoneEntry, buttons))
	window.Show()

}
