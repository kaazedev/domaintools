package lookup

import (
	"context"
	"fmt"
	"github.com/ammario/ipisp/v2"
	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/likexian/whois"
	whois_parser "github.com/likexian/whois-parser"
	"github.com/pearleascent/domaintools/internal/crt"
	"log"
	"net"
	"strings"
)

func NSLookup(domain string) {
	color.HiYellow("Searching for NSs...")
	ns, err := net.LookupNS(domain)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %s\n\n", err))
		return
	}

	for i, n := range ns {
		fmt.Printf("%s %s\n", color.HiGreenString("[%d]", i), color.HiWhiteString(n.Host))
	}

	fmt.Println(color.HiGreenString("Total NSs: %d\n\n", len(ns)))
}

func MXLookup(domain string) {
	color.HiYellow("Searching for MXs...")
	mx, err := net.LookupMX(domain)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %s\n\n", err))
		return
	}

	for i, n := range mx {
		fmt.Printf("%s %s\n", color.HiGreenString("[%d]", i), color.HiWhiteString(n.Host))
	}

	fmt.Println(color.HiGreenString("Total MXs: %d\n\n", len(mx)))
}

func CNAMELookup(domain string) {
	color.HiYellow("Searching for CNAME...")
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %s\n\n", err))
		return
	}

	fmt.Printf("%s %s\n", color.HiGreenString("[CNAME]"), color.HiWhiteString(cname))
	fmt.Println("")
}

func TXTLookup(domain string) {
	color.HiYellow("Searching for TXTs...")
	txt, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %s\n\n", err))
		return
	}

	for i, n := range txt {
		fmt.Printf("%s %s\n", color.HiGreenString("[%d]", i), color.HiWhiteString(n))
	}

	fmt.Println(color.HiGreenString("Total TXTs: %d\n\n", len(txt)))
}

func ALookup(domain string) {
	color.HiYellow("Searching for IPs...")
	ip, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %s\n\n", err))
		return
	}

	for i, n := range ip {
		resp, err := ipisp.LookupIP(context.Background(), net.ParseIP(n.String()))
		if err != nil {
			panic(err)
		}

		resp, err = ipisp.LookupASN(context.Background(), ipisp.ASN(resp.ASN))
		if err != nil {
			panic(err)
		}

		city, err := ipinfo.GetIPCity(n)
		if err != nil {
			log.Fatal(err)
		}

		country, err := ipinfo.GetIPCountry(n)
		if err != nil {
			log.Fatal(err)
		}

		location, err := ipinfo.GetIPLocation(n)
		if err != nil {
			log.Fatal(err)
		}

		organization, err := ipinfo.GetIPOrg(n)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s %s %s %s %s %s\n",
			color.HiGreenString("[%d]", i),
			color.HiWhiteString(n.String()),
			color.HiCyanString(country),
			color.HiCyanString(city),
			color.HiCyanString(location),
			color.HiGreenString(organization),
			color.HiGreenString(resp.ISPName),
		)
	}

	fmt.Println(color.HiGreenString("Total IPs: %d\n\n", len(ip)))
}

func Subdomains(name string) {
	color.HiYellow("Searching for subdomains...")
	subdomain, err := crt.RequestCRT(name)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %s\n\n", err))
		return
	}

	data := crt.RemoveDuplicate(subdomain)

	for i, sub := range data {
		sub = strings.Replace(sub, "\n", " ", -1)
		fmt.Printf("%s %s\n", color.HiGreenString("[%d]", i), color.HiWhiteString(sub))
	}

	fmt.Println(color.HiGreenString("Total subdomains: %d\n\n", len(data)))
}

func Whois(name string) {
	whoisRaw, err := whois.Whois(name)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %s\n\n", err))
		return
	}

	result, err := whois_parser.Parse(whoisRaw)
	if err != nil {
		fmt.Printf(color.HiRedString("Error: %s\n\n", err))
		return
	}

	fmt.Printf("%s %s\n", color.HiGreenString("[Registrar Name]"), color.HiWhiteString(result.Registrar.Name))
	fmt.Printf("%s %s\n", color.HiGreenString("[Registrar Email]"), color.HiWhiteString(result.Registrar.Email))
	fmt.Printf("%s %s\n", color.HiGreenString("[Registrar Phone]"), color.HiWhiteString(result.Registrar.Phone))
	fmt.Printf("%s %s\n", color.HiGreenString("[Registered At]"), color.HiWhiteString(result.Domain.CreatedDate))
	fmt.Printf("%s %s\n", color.HiGreenString("[Expiration At]"), color.HiWhiteString(result.Domain.ExpirationDate))
	fmt.Printf("%s %s\n", color.HiGreenString("[Status]"), color.HiWhiteString(result.Domain.Status[0]))
}
