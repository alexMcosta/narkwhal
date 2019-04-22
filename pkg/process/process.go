package process

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/alexmcosta/narkwhal/pkg/ingest"
)

// GetFlags does the work of getting the flags!
func GetFlags() (string, string, string) {
	accountFlag := flag.String("account", "default", "Lets you select witch AWS account you would like to make changes to")
	regionFlag := flag.String("region", "us-east-1", "Lets you select which region you would like to run Narkwhal on")
	timeFlag := flag.String("time", "0s", "Lets you select the amount of time a volume has been available based on MS, seconds, and Hours")
	flag.Parse()

	fmt.Printf("account: %s, region: %s, Not used within: %s\n", *accountFlag, *regionFlag, *timeFlag)
	return *accountFlag, *regionFlag, *timeFlag
}

// Confirm makes sure that the user either gives a yes or no answer
func Confirm() bool {
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadRune()

	switch text {
	case 'y', 'Y':
		return true
	case 'n', 'N':
		return false
	default:
		fmt.Println("Please type y or n and then press enter: ")
		return Confirm()
	}
}

func ListVolumesAndConfirm(account string, region string, time string) {

	if time == "0s" {
		fmt.Println("---------------------")
		ingest.ListVolumeIDs(account, region)
		fmt.Println("---------------------")
		fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")
		response := Confirm()
		if response == true {
			ingest.RemoveAvailableEBSNoTime(account, region)
		} else {
			fmt.Println("---------\nExiting: Nothing deleted\n---------")
		}
	} else {
		sliceOfVolumes := ingest.FilterOldVolumesByTime(account, region, time)
		fmt.Println("---------------------")
		for _, value := range sliceOfVolumes {
			fmt.Println(value)
		}
		fmt.Println("---------------------")
		fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")
		response := Confirm()
		if response == true {
			ingest.RemoveAvailableEBSYesTime(account, region, sliceOfVolumes)
		} else {
			fmt.Println("---------\nExiting: Nothing deleted\n---------")
		}
	}
}
