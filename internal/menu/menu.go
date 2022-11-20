package menu

import (
	"fmt"
	"github.com/fatih/color"
)

func Menu() (int, string) {
	var action int
	var ipordomain string

	color.HiGreen("Welcome to the DomainTools")
	color.HiWhite("Please select an action:")

	fmt.Println("")

	color.HiWhite("Domain Records:")
	fmt.Printf("%s %s\n", color.HiGreenString("[1]"), color.HiWhiteString("NS Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[2]"), color.HiWhiteString("MX Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[3]"), color.HiWhiteString("A Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[4]"), color.HiWhiteString("TXT Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[5]"), color.HiWhiteString("CNAME Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[6]"), color.HiWhiteString("Subdomains"))
	fmt.Printf("%s %s\n", color.HiGreenString("[7]"), color.HiWhiteString("All domain information"))

	color.HiWhite("\nWhois Service:")
	fmt.Printf("%s %s\n", color.HiGreenString("[8]"), color.HiWhiteString("Whois"))

	color.HiWhite("\nCheck the availability to register of domain:")
	fmt.Printf("%s %s\n", color.HiGreenString("[9]"), color.HiWhiteString("Single"))
	fmt.Printf("%s %s\n", color.HiGreenString("[10]"), color.HiWhiteString("Bulk"))
	fmt.Println("")

	fmt.Printf("%s\n", color.HiYellowString("Select an action:"))
	fmt.Scanln(&action)

	if action == 10 {
		fmt.Printf("%s\n", color.HiYellowString("Enter an domain without TLD (like \"domainname\"):"))
	} else {
		fmt.Printf("%s\n", color.HiYellowString("Enter an domain:"))
	}
	fmt.Scanln(&ipordomain)

	fmt.Println("")

	return action, ipordomain
}
