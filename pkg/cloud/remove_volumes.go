package cloud

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// RemoveAvailableVolumes Takes a slice of Volume ID's and deletes them
func RemoveVolumes(acctData map[string]map[string][]string) {

	for acct, regions := range acctData {
		for region, volID := range regions {

			svc := createEC2Session(acct, region)

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

}
