package process

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/alexmcosta/narkwhal/pkg/ingest"
	"github.com/aws/aws-sdk-go/service/ec2"
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

// GetSliceOfIDs takes the struct of EBS volume data and retuens a slice of only the ID's
func GetSliceOfIDs(volume *ec2.DescribeVolumesOutput) []string {

	var sliceOfIDs []string

	for _, value := range volume.Volumes {
		sliceOfIDs = append(sliceOfIDs, *value.VolumeId)
	}

	return sliceOfIDs
}

// ListVolumesAndConfirm takes a list of VolumeID's and confirms if the user wants them removed
func ListVolumesAndConfirm(filteredSliceOfVolumes []string, account string, region string, time string) {

	if filteredSliceOfVolumes == nil {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("EXITING: There are no available EBS volumes to remove in the %s region of account %s that have been available for at least %s\n", region, account, time)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		os.Exit(1)
	}

	fmt.Println("---------------------")
	for _, value := range filteredSliceOfVolumes {
		fmt.Println(value)
	}
	fmt.Println("---------------------")
	fmt.Println("Would you like to remove the above EBS Volumes? (y/n): ")
	response := confirm()
	if response == true {
		ingest.RemoveAvailableVolumes(filteredSliceOfVolumes, account, region)
	} else {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("EXITING: Nothing Deleted")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	}
}
