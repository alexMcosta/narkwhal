package ingest

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// RemoveAvailableVolumes Takes a slice of Volume ID's and deletes them
func RemoveAvailableVolumes(input []string, accountFlag string, regionFlag string) {

	for _, value := range input {

		svc := createEC2Session(accountFlag, regionFlag)

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
