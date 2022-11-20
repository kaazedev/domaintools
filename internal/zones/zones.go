package zones

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/haccer/available"
	"time"
)

func Single(domain string) {
	if available.Domain(domain) {
		fmt.Printf("%s %s \n", color.HiGreenString("[+]"), color.HiWhiteString("%s", domain))
	} else {
		fmt.Printf("%s %s \n", color.HiRedString("[X]"), color.HiWhiteString("%s", domain))
	}
}

func Multiple(domain string) {
	FromFile()

	for _, tld := range TLDList.List {
		Single(domain + "." + tld)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("")
}
