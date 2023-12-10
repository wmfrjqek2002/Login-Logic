package Logic

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
	"strings"
)

type info struct {
	ID       string
	Password string
	Name     string
	Phone    string
	Admin    Adminstrator `json:",omitempty"`
}

type Adminstrator struct {
	Privilege bool
}

var Information []info
var TempData []info

func SignIn(a fyne.App) {

	SignIn_W := a.NewWindow("회원가입창")
	SignIn_W.Resize(fyne.NewSize(300, 400))
	a.Settings().SetTheme(theme.DarkTheme())

	welcome := canvas.NewText("회원가입 페이지입니다.", color.White)
	welcome.TextSize = 30
	welcome.Move(fyne.NewPos(5, 5))

	spacer := widget.NewLabel("")

	errMsg := canvas.NewText("", color.NRGBA{255, 0, 0, 255})

	ID_Label := canvas.NewText("아이디", color.White)
	User_Id := widget.NewEntry()

	PS_Label := canvas.NewText("비밀번호", color.White)
	User_PS := widget.NewPasswordEntry()

	Name_Label := canvas.NewText("이름", color.White)
	User_Name := widget.NewEntry()

	Phone_Label := canvas.NewText("휴대폰", color.White)
	User_Phone := widget.NewEntry()

	Submit_Button := widget.NewButton("회원가입", func() {
		id_len := len(User_Id.Text)
		ps_len := len(User_PS.Text)
		name_len := len(User_Name.Text)
		phone_len := len(User_Phone.Text)

		newID := User_Id.Text
		newPhone := User_Phone.Text
		idExists := false
		for _, existingData := range TempData {
			if existingData.ID == newID {
				idExists = true
				break
			}
		}
		phoneExists := false
		for _, existingData := range TempData {
			if existingData.Phone == newPhone {
				phoneExists = true
				break
			}
		}

		if id_len > 5 && ps_len >= 8 && name_len > 0 && phone_len >= 10 {
			if idExists {
				errMsg.Text = "이미 존재하는 아이디입니다."
				errMsg.Refresh()
			} else if phoneExists {
				errMsg.Text = "이미 존재하는 휴대폰 번호입니다."
				errMsg.Refresh()
			} else {
				NewData := &info{
					ID:       User_Id.Text,
					Password: User_PS.Text,
					Name:     User_Name.Text,
					Phone:    User_Phone.Text,
				}

				if IsAdministrator(User_Id.Text) {
					NewData.Admin = Adminstrator{
						Privilege: true,
					}
				}

				Information = append(Information, *NewData)
				Information = append(TempData, *NewData)

				info_data, _ := json.MarshalIndent(Information, "", " ")
				UserCreateFile := fmt.Sprintf("User\\%s.txt", User_Id.Text)
				os.WriteFile(UserCreateFile, info_data, 0644)

				User_Id.Text = ""
				User_PS.Text = ""
				User_Name.Text = ""
				User_Phone.Text = ""

				User_Id.Refresh()
				User_PS.Refresh()
				User_Name.Refresh()
				User_Phone.Refresh()

				complete := a.NewWindow("complete")
				cLabel := canvas.NewText("회원가입이 완료되었습니다.", color.White)
				cButton := widget.NewButton("확인", func() {
					complete.Close()
					SignIn_W.Close()
				})
				complete.SetContent(container.NewVBox(cLabel, cButton))
				complete.Show()
			}

		} else if id_len < 6 {
			errMsg.Text = "아이디를 6자 이상 입력하시오"
			errMsg.Refresh()
		} else if ps_len < 9 {
			errMsg.Text = "비밀번호를 8자 이상 입력하시오"
			errMsg.Refresh()
		} else if name_len < 1 {
			errMsg.Text = "이름을 입력하세요."
			errMsg.Refresh()
		} else if phone_len < 10 {
			errMsg.Text = "전화번호를 제대로 입력하세요."
			errMsg.Refresh()
		}

	})

	Cancel_Button := widget.NewButton("취소", func() {
		SignIn_W.Close()
	})

	SignIn_W.SetContent(container.NewVBox(container.NewWithoutLayout(welcome), spacer,
		ID_Label, User_Id, PS_Label, User_PS, Name_Label, User_Name, Phone_Label, User_Phone, errMsg,
		Submit_Button, Cancel_Button))
	SignIn_W.Show()
}

func IsAdministrator(userId string) bool {
	return strings.HasPrefix(userId, "admin")
}
