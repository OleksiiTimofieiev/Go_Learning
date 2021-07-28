package main

import (
	"fmt"
	"strings"

	flag "github.com/spf13/pflag"
)

// Define the variables corresponding to the command line parameters
var cliName = flag.StringP("name", "n", "nick", "Input Your Name")
var cliAge = flag.IntP("age", "a", 22, "Input Your Age")
var cliGender = flag.StringP("gender", "g", "male", "Input Your Gender")
var cliOK = flag.BoolP("ok", "o", false, "Input Are You OK")
var cliDes = flag.StringP("des-detail", "d", "", "Input Description")
var cliOldFlag = flag.StringP("badflag", "b", "just for test", "Input badflag")

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func main() {
	// Function to set standardized parameter names
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)

	// Set NoOptDefVal for age parameter
	flag.Lookup("age").NoOptDefVal = "25"

	// Mark the badflag parameter as being obsolete, please use the des-detail parameter
	flag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")

	// Mark the shorthand of the badflag parameter as being obsolete, please use the shorthand parameter of des-detail
	flag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")

	// Hide the parameter gender in the help document
	flag.CommandLine.MarkHidden("badflag")

	// parse the command line parameters passed by the user into the value of the corresponding variable
	flag.Parse()

	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("ok=", *cliOK)
	fmt.Println("des=", *cliDes)

}
