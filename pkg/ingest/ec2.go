package ingest

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GrabAvailableVolumes Uses the AWS SDK to search for all available volumes in the current region
func GrabAvailableVolumes(accountFlag string, regionFlag string) (volume *ec2.DescribeVolumesOutput) {

	svc := createSession(accountFlag, regionFlag)
	// Let us filter for all available EBS volumes
	input := &ec2.DescribeVolumesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("status"),
				Values: []*string{
					aws.String("available"),
				},
			},
		},
	}

	// Go get them volumes and send an AWS error if there is one
	volumes, err := svc.DescribeVolumes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
				os.Exit(1)
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	// If there are no available EBS volumes then quit application
	if volumes == nil {
		fmt.Println("grabAvailableVolumes: There are no available EBS volumes")
		os.Exit(1)
		return
	}

	return volumes
}

// ListVolumeIDs will list the volume IDs that are about to be deleted
func ListVolumeIDs(accountFlag string, regionFlag string) {
	input := GrabAvailableVolumes(accountFlag, regionFlag)

	if input.Volumes == nil {
		fmt.Println("~~~~~~~~~~~~~~~~~~()~~~~")
		fmt.Printf("EXITING: There are no available EBS volumes in the %s region to remove\n", regionFlag)
		fmt.Println("~~~~~~~~~~~~~~~~~~()~~~~")
		fmt.Println("---------------------")
		os.Exit(1)
	} else {
		for _, value := range input.Volumes {
			fmt.Println(*value.VolumeId)
		}
	}
}

// TODO: Refactor: These Remove functions can either be combined or chopped down

// RemoveAvailableEBS Removes all avail able EBS volumes based on the current default region
func RemoveAvailableEBSNoTime(accountFlag string, regionFlag string) {

	input := GrabAvailableVolumes(accountFlag, regionFlag)

	for _, value := range input.Volumes {

		svc := createSession(accountFlag, regionFlag)

		deleteInput := &ec2.DeleteVolumeInput{
			VolumeId: aws.String(*value.VolumeId),
		}

		_, err := svc.DeleteVolume(deleteInput)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
		}

		fmt.Println("Successfully removed", *value.VolumeId)

	}
}

func RemoveAvailableEBSYesTime(accountFlag string, regionFlag string, input []string) {

	for _, value := range input {

		svc := createSession(accountFlag, regionFlag)

		deleteInput := &ec2.DeleteVolumeInput{
			VolumeId: aws.String(value),
		}

		_, err := svc.DeleteVolume(deleteInput)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
		}

		fmt.Println("Successfully removed", value)

	}

}
