package ingest

import (
	"fmt"

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

// Go get all the "available" / unused volumes snapshot ID's
func RemoveOldEBS(input *ec2.DeleteVolumeInput) {

	svc := createSession()

	result, err := svc.DeleteVolume(input)
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

	fmt.Println(result)
}

// Go get
