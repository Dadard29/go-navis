package common

import "fmt"

func PrintResponse(output string) {
	fmt.Println("[+] " + output)
}

func PrintError(error string) {
	fmt.Println("[!] " + error)
}
