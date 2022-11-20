package main

import (
	"fmt"
	"github.com/haccer/available"
	"github.com/pearleascent/domaintools/internal/lookup"
	"github.com/pearleascent/domaintools/internal/menu"
)

func main() {
	action, domain := menu.Menu()

	switch action {
	case 1:
		lookup.NSLookup(domain)
		break
	case 2:
		lookup.MXLookup(domain)
		break
	case 3:
		lookup.ALookup(domain)
		break
	case 4:
		lookup.TXTLookup(domain)
		break
	case 5:
		lookup.CNAMELookup(domain)
		break
	case 6:
		lookup.Subdomains(domain)
		break
	case 7:
		lookup.NSLookup(domain)
		lookup.MXLookup(domain)
		lookup.ALookup(domain)
		lookup.TXTLookup(domain)
		lookup.CNAMELookup(domain)
		lookup.Subdomains(domain)
		break
	case 8:
		if available.Domain(domain) {
			fmt.Printf("Available for registration\n")
		} else {
			fmt.Printf("Not available for registration\n")
		}
		break
	}

	fmt.Scanln()
}
