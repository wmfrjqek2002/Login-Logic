package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"market/Logic"
	"market/mainApp"
	"os"
)

type Data struct {
	ID       string
	Password string
}

func main() {

	LoginOk := false
	IdErr := true
	PasswordErr := false

	err := os.Setenv("FYNE_FONT", "data\\NanumSquareRoundL.ttf")
	if err != nil {
		fmt.Println(err)
	}
	a := app.New()
	w := a.NewWindow("로그인 프로그램")
	w.Resize(fyne.NewSize(320, 220))
	a.Settings().SetTheme(theme.DarkTheme())

	login_logo := canvas.NewImageFromFile("data\\login_logo.png")
	login_logo.Resize(fyne.NewSize(100, 100))

	error_msg := canvas.NewText("", color.NRGBA{255, 0, 0, 255})
	error_msg.Move(fyne.NewPos(130, 100))

	ID := widget.NewFormItem("아이디", widget.NewEntry())
	Password := widget.NewFormItem("비밀번호", widget.NewPasswordEntry())
	login_form := widget.NewForm(
		ID,
		Password,
	)
	login_form.Resize(fyne.NewSize(200, 100))
	login_form.Move(fyne.NewPos(110, 15))

	login_Button := widget.NewButton("로그인", func() {
		LoginId := ID.Widget.(*widget.Entry).Text
		LoginPassword := Password.Widget.(*widget.Entry).Text

		User := fmt.Sprintf("User\\%s.txt", LoginId)
		var data []Data
		data_file, _ := os.ReadFile(User)
		json.Unmarshal(data_file, &data)

		for _, login := range data {
			if login.ID == LoginId && login.Password == LoginPassword {
				LoginOk = true
			}
		}

		for _, login := range data {
			if login.ID == LoginId {
				IdErr = false
				break
			}
		}

		for _, login := range data {
			if login.Password != LoginPassword {
				PasswordErr = true
				break
			}
		}

		if LoginOk {
			mainApp.App(a, w, LoginId)
			ID.Widget.(*widget.Entry).SetText("")
			Password.Widget.(*widget.Entry).SetText("")
			error_msg.Text = ""
			error_msg.Refresh()
			w.Hide()
		} else if PasswordErr {
			error_msg.Text = "비밀번호가 일치하지 않습니다."
			error_msg.Refresh()
		} else if IdErr {
			error_msg.Text = "존재하지 않는 아이디입니다."
			error_msg.Refresh()
		}
	})
	login_Button.Resize(fyne.NewSize(200, 30))
	login_Button.Move(fyne.NewPos(60, 130))

	SignIn_Button := widget.NewButton("회원가입", func() {
		go Logic.SignIn(a)
	})
	SignIn_Button.Resize(fyne.NewSize(200, 30))
	SignIn_Button.Move(fyne.NewPos(60, 170))

	w.SetContent(container.NewHBox(container.NewWithoutLayout(login_logo, login_form, error_msg, login_Button, SignIn_Button)))
	w.ShowAndRun()
}
