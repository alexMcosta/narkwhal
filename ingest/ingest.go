package ingest

import (
	"fmt"
	"os"

	"github.com/alexmcosta/narkwhal/process"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// createSession Creates gthe sessions needed to work with the AWS SDk
func createSession() *ec2.EC2 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ec2.New(sess)
	return svc
}

// grabAvailableVolumeIDs Uses the AWS SDK to search for all available volumes in the current region
func GrabAvailableVolumesIDs() (volume *ec2.DescribeVolumesOutput) {

	if volume == nil {
		fmt.Println("GrabAvailableVolumesIDs: There are no available EBS volumes")
		os.Exit(1)
		return
	} else {

		// Create a session
		svc := createSession()
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

		volumes, err := svc.DescribeVolumes(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
			return
		}
		return volumes
	}
}

// RemoveAvailableEBS Removes all avail able EBS volumes based on the current default region
func RemoveAvailableEBS(input []process.Volumes) {
	//Print a message if the
	if input == nil {
		fmt.Println("RemoveAvailableEBSVolumes(): There are no EBS volumes to remove")
	} else {

		for _, value := range input {

			svc := createSession()

			deleteInput := &ec2.DeleteVolumeInput{
				VolumeId: aws.String(value.VolumeId),
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
