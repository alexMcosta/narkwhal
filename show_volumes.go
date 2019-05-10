package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexmcosta/narkwhal/pkg/cloud"
)

// showVolumes takes a list of volume ID's and confirms if the user wants them removed
// If the user wants them removed it calls on the cloud package
func showVolumes(acctMap map[string]map[string][]string, acc []string, time string) {

	remVol := make(map[string][]string)

	for aa, rr := range acctMap {

		for r, idSlice := range rr {

			if idSlice == nil {
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
				fmt.Printf("No volumes meet specified criteria in %s, %s\n", aa, r)
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
			} else {
				fmt.Println("~~~~~~~~~~~")
				fmt.Printf("Account: %s --- Region: %s\n", aa, r)
				fmt.Println("~~~~~~~~~~~")

				for _, volID := range idSlice {
					fmt.Println(volID)
				}

				fmt.Println("---------------------")
				fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")
				response := confirm()
				if response == true {
					remVol[r] = idSlice
				}
			}
		}
	}
	if remVol != nil {
		cloud.RemoveVolumes(acc, remVol)
	}
}

// confirm makes sure that the user either gives a yes or no answer
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
