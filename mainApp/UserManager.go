package mainApp

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"path/filepath"
	"strings"
)

func UserManager(a fyne.App) {
	window := a.NewWindow("사용자관리")
	window.Resize(fyne.NewSize(300, 300))

	userDirectory := "User"
	fileNames, err := listTextFiles(userDirectory)

	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	list := widget.NewList(
		func() int {
			return len(fileNames)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(i int, item fyne.CanvasObject) {
			if label, ok := item.(*widget.Label); ok {
				filemultxt := strings.TrimSuffix(fileNames[i], ".txt")
				label.SetText(filemultxt)
			}
		})

	list.OnSelected = func(id int) {
		selectedFileName := fileNames[id]
		jsonData, err := readJSONFromFile(filepath.Join(userDirectory, selectedFileName))
		if err != nil {
			fmt.Println("파일 읽기 오류:", err)
			return
		}

		showData(a, selectedFileName, jsonData)
	}
	list.Resize(fyne.NewSize(300, 200))

	content := container.NewVBox(
		container.NewWithoutLayout(list),
	)

	window.SetContent(content)
	window.Show()
}

func listTextFiles(directory string) ([]string, error) {
	var fileNames []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			fileNames = append(fileNames, info.Name())
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return fileNames, nil
}

func readJSONFromFile(filePath string) ([]Data, error) {
	var data []Data
	data_file, err := os.ReadFile(filePath)

	json.Unmarshal(data_file, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
