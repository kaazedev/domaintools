package main

import (
	"fmt"
	"github.com/pearleascent/domaintools/internal/lookup"
	"github.com/pearleascent/domaintools/internal/menu"
	"github.com/pearleascent/domaintools/internal/zones"
	"github.com/pearleascent/domaintools/pgk/clearconsole"
)

func main() {
	for {
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
			lookup.Whois(domain)
			break
		case 9:
			zones.Single(domain)
			break
		case 10:
			zones.Multiple(domain)
			break
		}

		fmt.Println("\nPress any key to start new round...")
		fmt.Scanln()

		clearconsole.Clear()
	}
}
