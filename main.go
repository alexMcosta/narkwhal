package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexmcosta/narkwhal/pkg/ingest"
	"github.com/alexmcosta/narkwhal/pkg/process"
)

// A looping confirmation function to make sure the user selects yes or no
func confirm() bool {
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadRune()

	switch text {
	case 'y', 'Y':
		return true
	case 'n', 'N':
		return false
	default:
		fmt.Println("Please type y or n and then press enter: ")
		return confirm()
	}
}

func main() {
	// Flags
	account, region := process.GetFlags()

	//Tell user the volume ID's and confirm deletion
	fmt.Println("---------------------")
	ingest.ListVolumeIDs(account, region)
	fmt.Println("---------------------")
	fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")

	// Get confirmation they want the EBS volumes deleted
	response := confirm()
	if response == true {
		ingest.RemoveAvailableEBS(account, region)
	} else {
		fmt.Println("---------\nExiting: Nothing deleted\n---------")
	}
}
