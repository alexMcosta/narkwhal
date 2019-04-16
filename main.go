package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/alexmcosta/narkwhal/ingest"
	"github.com/alexmcosta/narkwhal/process"
)

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
	data, _ := json.Marshal(volumeIDs)

	// Process the ID's And make them structs
	volumes := process.GetStructOfVolumes(data)

	//Tell user the volume ID's and confirm deletion
	fmt.Println("---------------------")
	fmt.Println(volumes)
	fmt.Println("---------------------")
	fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")

	response := confirm()

	if response == true {
		ingest.RemoveAvailableEBS(volumes, *accountFlag, *regionFlag)
	} else {
		fmt.Println("---------\nExiting: Nothing deleted\n---------")
	}

	// Remove any availble structs
	//ingest.RemoveAvailableEBS(volumes, *accountFlag, *regionFlag)
}
