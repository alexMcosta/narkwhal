package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/alexmcosta/narkwhal/ingest"
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
	accountFlag := flag.String("account", "default", "Lets you select witch AWS account you would like to make changes to")
	regionFlag := flag.String("region", "us-east-1", "Lets you select which region you would like to run Narkwhal on")
	flag.Parse()
	fmt.Printf("account: %s, region: %s\n", *accountFlag, *regionFlag)

	// Grab all available volumes and marshal them
	volumeIDs := ingest.GrabAvailableVolumesIDs(*accountFlag, *regionFlag)
	// data, _ := json.Marshal(volumeIDs)

	//fmt.Println(*volumeIDs.Volumes[0].VolumeId)

	// Process the data and make structs
	// volumes := process.GetStructOfVolumes(data)

	//Tell user the volume ID's and confirm deletion
	fmt.Println("---------------------")
	fmt.Println(*volumeIDs.Volumes[0].VolumeId)
	fmt.Println("---------------------")
	fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")

	// Get confirmation they want the EBS volumes deleted
	response := confirm()
	if response == true {
		ingest.RemoveAvailableEBS(*accountFlag, *regionFlag)
	} else {
		fmt.Println("---------\nExiting: Nothing deleted\n---------")
	}
}
