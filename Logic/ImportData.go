package Logic

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	ID       string
	Password string
	Name     string
	Phone    string
	Admin    struct {
		Privilege bool `json:"Privilege"`
	} `json:"Admin"`
}

func ImportData(User string) []Data {
	var data []Data
	data_file, _ := os.ReadFile(fmt.Sprintf("User\\%s.txt", User))
	json.Unmarshal(data_file, &data)

	return data
}
