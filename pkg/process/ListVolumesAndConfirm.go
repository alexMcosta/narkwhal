package process

import (
	"fmt"

	"github.com/alexmcosta/narkwhal/pkg/ingest"
)

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
