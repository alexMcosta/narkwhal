package ingest

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// createSession Creates the sessions needed to work with the AWS SDk
func createSession(accountFlag string, regionFlag string) *ec2.EC2 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(regionFlag)},
		Profile: accountFlag,
	}))
	svc := ec2.New(sess)
	return svc
}

// GrabAvailableVolumeIDs Uses the AWS SDK to search for all available volumes in the current region
func GrabAvailableVolumesIDs(accountFlag string, regionFlag string) (volume *ec2.DescribeVolumesOutput) {

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
	// Else continue
	// if volume == nil {
	// 	fmt.Println("GrabAvailableVolumesIDs: There are no available EBS volumes")
	// 	os.Exit(1)
	// 	return
	// } else {
	return volumes
	// }
}

// RemoveAvailableEBS Removes all avail able EBS volumes based on the current default region
func RemoveAvailableEBS(accountFlag string, regionFlag string) {
	input := GrabAvailableVolumesIDs(accountFlag, regionFlag)

	// Print a message if there are no EBS volumes to delete
	// Or keep going to remove the volumes
	if input == nil {
		fmt.Println("RemoveAvailableEBSVolumes(): There are no EBS volumes to remove")
	} else {

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

			fmt.Println("Successfully removed", value.VolumeId)

		}
	}
}
