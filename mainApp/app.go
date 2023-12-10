package mainApp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Data struct {
	ID       string
	Password string
	Name     string
	Phone    string
}

func App(a fyne.App, w fyne.Window, UserID string) {

	window := a.NewWindow("hello")
	window.Resize(fyne.NewSize(600, 500))

	welcome := canvas.NewText("환영합니다!", color.White)
	welcome.TextSize = 60
	_ = widget.NewButton("사용자 정보 관리", func() {

	})
	User_Manager := widget.NewButton("사용자 관리", func() {

	})
	User_Manager.Resize(fyne.NewSize(100, 30))
	User_Manager.Move(fyne.NewPos(390, 0))

	info_Modic := widget.NewButton("정보수정", func() {
		InfoModict(a, UserID)
	})
	info_Modic.Resize(fyne.NewSize(100, 30))
	info_Modic.Move(fyne.NewPos(495, 0))

	logout := widget.NewButton("로그아웃", func() {
		window.Close()
		w.Show()
	})
	logout.Resize(fyne.NewSize(100, 30))
	logout.Move(fyne.NewPos(0, 0))

	close := widget.NewButton("닫기", func() {
		window.Close()
		w.Close()

	})
	close.Resize(fyne.NewSize(100, 30))
	close.Move(fyne.NewPos(105, 0))

	menu := container.NewWithoutLayout(logout, close, User_Manager, info_Modic)
	window.SetContent(container.NewVBox(menu, container.NewCenter(welcome)))
	window.Show()
}
