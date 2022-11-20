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

	color.HiWhite("Domain information: ")
	fmt.Printf("%s %s\n", color.HiGreenString("[1]"), color.HiWhiteString("NS Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[2]"), color.HiWhiteString("MX Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[3]"), color.HiWhiteString("A Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[4]"), color.HiWhiteString("TXT Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[5]"), color.HiWhiteString("CNAME Lookup"))
	fmt.Printf("%s %s\n", color.HiGreenString("[6]"), color.HiWhiteString("Subdomains"))
	fmt.Printf("%s %s\n", color.HiGreenString("[7]"), color.HiWhiteString("All domain information"))
	fmt.Println("")
	color.HiWhite("Check the availability of domains:")
	fmt.Printf("%s %s\n", color.HiGreenString("[8]"), color.HiWhiteString("Single"))
	/*	fmt.Printf("%s %s\n", color.HiGreenString("[9]"), color.HiWhiteString("Bulk"))*/
	fmt.Println("")

	fmt.Printf("%s\n", color.HiGreenString("Select an action:"))
	fmt.Scanln(&action)

	//if action == 9 {
	//	fmt.Printf("%s\n", color.HiGreenString("Enter an domain without TDL (like \"domainname\"):"))
	//} else {
	//	fmt.Printf("%s\n", color.HiGreenString("Enter an domain:"))
	//}

	fmt.Printf("%s\n", color.HiGreenString("Enter an domain:"))
	fmt.Scanln(&ipordomain)

	fmt.Println("")

	return action, ipordomain
}
