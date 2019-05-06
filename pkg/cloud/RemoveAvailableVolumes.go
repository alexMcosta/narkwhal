package cloud

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// RemoveAvailableVolumes Takes a slice of Volume ID's and deletes them
func RemoveAvailableVolumes(accountFlag string, regionsData map[string][]string) {

	for region, volID := range regionsData {

		svc := createEC2Session(accountFlag, region)

		for _, volumeID := range volID {

			deleteInput := &ec2.DeleteVolumeInput{
				VolumeId: aws.String(volumeID),
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
			fmt.Printf("Successfully removed %s\n", volumeID)
		}

	}

}
