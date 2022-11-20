package zones

import (
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

type TLDs struct {
	List []string `json:"tlds"`
}

var TLDList TLDs

func FromFile() {
	jsonFile, err := os.Open("tlds.json")
	if err != nil {
		color.Red("Error opening tlds.json")
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		color.Red("Error reading tlds.json")
	}

	json.Unmarshal(byteValue, &TLDList)
	color.Green("Loaded %d TLDs", len(TLDList.List))
}
