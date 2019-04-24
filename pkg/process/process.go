package process

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alexmcosta/narkwhal/pkg/ingest"
)

// MultiRegion will sort through comma seperated regions and make a slice of strings
func MultiRegion(regions string) []string {
	var sliceOfRegions []string

	if strings.Contains(regions, ",") {
		sliceOfRegions := strings.Split(regions, ",")
		return sliceOfRegions
	} else {
		sliceOfRegions = append(sliceOfRegions, regions)
	}

	return sliceOfRegions

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

// ListVolumesAndConfirm takes a list of VolumeID's and confirms if the user wants them removed
func ListVolumesAndConfirm(filteredSliceOfVolumes map[string][]string, account string, time string) {

	removableVolumes := make(map[string][]string)

	for region, sliceOfIDs := range filteredSliceOfVolumes {

		if sliceOfIDs == nil {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Printf("No Available Volumes in %s\n", region)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		} else {
			fmt.Println("~~~~~~~~~~~")
			fmt.Println(region)
			fmt.Println("~~~~~~~~~~~")

			for _, volumeID := range sliceOfIDs {
				fmt.Println(volumeID)
			}
			fmt.Println("---------------------")
			fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")
			response := confirm()
			if response == true {
				removableVolumes[region] = sliceOfIDs
			}
		}
	}
	if removableVolumes != nil {
		ingest.RemoveAvailableVolumes(account, removableVolumes)
	}
}
