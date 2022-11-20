package lookup

import (
	"context"
	"fmt"
	"github.com/ammario/ipisp/v2"
	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/pearleascent/domaintools/internal/crt"
	"log"
	"net"
	"strings"
)

func NSLookup(domain string) {
	color.HiYellow("Searching for NSs...")
	ns, err := net.LookupNS(domain)
	if err != nil {
		panic(err)
	}

	for i, n := range ns {
		fmt.Printf("%s %s\n", color.HiGreenString("[%d]", i), color.HiWhiteString(n.Host))
	}

	fmt.Println(color.HiGreenString("Total NSs: %d\n", len(ns)))
}

func MXLookup(domain string) {
	color.HiYellow("Searching for MXs...")
	mx, err := net.LookupMX(domain)
	if err != nil {
		panic(err)
	}

	for i, n := range mx {
		fmt.Printf("%s %s\n", color.HiGreenString("[%d]", i), color.HiWhiteString(n.Host))
	}

	fmt.Println(color.HiGreenString("Total MXs: %d\n", len(mx)))
}

func CNAMELookup(domain string) {
	color.HiYellow("Searching for CNAME...")
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s\n", color.HiGreenString("[CNAME]"), color.HiWhiteString(cname))
	fmt.Println("")
}

func TXTLookup(domain string) {
	color.HiYellow("Searching for TXTs...")
	txt, err := net.LookupTXT(domain)
	if err != nil {
		panic(err)
	}

	for i, n := range txt {
		fmt.Printf("%s %s\n", color.HiGreenString("[%d]", i), color.HiWhiteString(n))
	}

	fmt.Println(color.HiGreenString("Total TXTs: %d\n", len(txt)))
}

func ALookup(domain string) {
	color.HiYellow("Searching for IPs...")
	ip, err := net.LookupIP(domain)
	if err != nil {
		panic(err)
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

	fmt.Println(color.HiGreenString("Total IPs: %d\n", len(ip)))
}

func Subdomains(name string) {
	color.HiYellow("Searching for subdomains...")
	subdomain, err := crt.RequestCRT(name)
	if err != nil {
		panic(err)
	}

	data := crt.RemoveDuplicate(subdomain)

	for i, sub := range data {
		sub = strings.Replace(sub, "\n", " ", -1)
		fmt.Printf("%s %s\n", color.HiGreenString("[%d]", i), color.HiWhiteString(sub))
	}

	fmt.Println(color.HiGreenString("Total subdomains: %d\n", len(data)))
}
