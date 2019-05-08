package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexmcosta/narkwhal/pkg/cloud"
)

// showVolumes takes a list of volume ID's and confirms if the user wants them removed
// If the user wants them removed it calls on the cloud package
func showVolumes(volSlice map[string][]string, acc string, time string) {

	remVol := make(map[string][]string)

	for reg, idSlice := range volSlice {

		if idSlice == nil {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Printf("No volumes meet specified criteria in %s\n", reg)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		} else {
			fmt.Println("~~~~~~~~~~~")
			fmt.Println(reg)
			fmt.Println("~~~~~~~~~~~")

			for _, volID := range idSlice {
				fmt.Println(volID)
			}

			fmt.Println("---------------------")
			fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")
			response := confirm()
			if response == true {
				remVol[reg] = idSlice
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
