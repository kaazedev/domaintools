package crt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CRTData struct {
	CommonName string `json:"common_name"`
	NameValue  string `json:"name_value"`
}

func RemoveDuplicate(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func RequestCRT(domain string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	sb := []byte(body)
	var subdomains []CRTData
	err = json.Unmarshal(sb, &subdomains)
	if err != nil {
		panic(err)
	}

	output := make([]string, 0)
	for _, subdomains := range subdomains {
		output = append(output, subdomains.CommonName)
		output = append(output, subdomains.NameValue)
	}

	return output, nil

}
