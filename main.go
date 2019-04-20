package main

import (
	"fmt"

	"github.com/alexmcosta/narkwhal/pkg/ingest"
	"github.com/alexmcosta/narkwhal/pkg/process"
)

func main() {
	// Flags
	account, region := process.GetFlags()

	fmt.Println(ingest.GrabAvailableVolumes(account, region))

	//Tell user the volume ID's and confirm deletion
	fmt.Println("---------------------")
	ingest.ListVolumeIDs(account, region)
	fmt.Println("---------------------")
	fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")

	// Get confirmation they want the EBS volumes deleted
	response := process.Confirm()
	if response == true {
		ingest.RemoveAvailableEBS(account, region)
	} else {
		fmt.Println("---------\nExiting: Nothing deleted\n---------")
	}
}
