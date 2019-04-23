package process

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alexmcosta/narkwhal/pkg/ingest"
)

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

	for region, sliceOfIDs := range filteredSliceOfVolumes {

		if sliceOfIDs == nil {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Printf("No Available Volumes in %s\n", region)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		} else {
			fmt.Println("~~~~~~~~~~~")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(region)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("~~~~~~~~~~~")

			for _, volumeID := range sliceOfIDs {
				fmt.Println(volumeID)
			}
			fmt.Println("---------------------")
			fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")
			response := confirm()
			if response == true {
				ingest.RemoveAvailableVolumes(account, filteredSliceOfVolumes)
			} else {
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
				fmt.Println("EXITING: Nothing Deleted")
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
			}
		}

	}
	// if filteredSliceOfVolumes == nil {
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	// 	fmt.Printf("EXITING: There are no available EBS volumes to remove with the specified fields")
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	// 	os.Exit(1)
	// }

	// fmt.Println("---------------------")
	// for _, value := range filteredSliceOfVolumes {
	// 	fmt.Println(value)
	// }
	// fmt.Println("---------------------")
	// fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")
	// response := confirm()
	// if response == true {
	// 	ingest.RemoveAvailableVolumes(filteredSliceOfVolumes, account, regions)
	// } else {
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	// 	fmt.Println("EXITING: Nothing Deleted")
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	// }
}
