package cloud

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GrabAvailableVolumes Uses the AWS SDK to search for all available volumes in the specified regions
func GrabAvailableVolumes(accountFlag string, regions []string) map[string][]string {

	// Make a map of regions so that way we can have the regions be the key with a slice of volume IDs
	mapOfRegions := make(map[string][]string)

	for _, region := range regions {

		svc := createEC2Session(accountFlag, region)

		// Struct to search for volumes with Status: Available
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
		}

		mapOfRegions[region] = getSliceOfIDs(volumes)

	}

	return mapOfRegions
}

// getSliceOfIDs takes the struct of EBS volume data and returns a slice of only the ID's
func getSliceOfIDs(volume *ec2.DescribeVolumesOutput) []string {

	var sliceOfIDs []string

	for _, value := range volume.Volumes {
		sliceOfIDs = append(sliceOfIDs, *value.VolumeId)
	}

	return sliceOfIDs
}
