package digest

import (
  "fmt"

  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/aws/aws-sdk-go/aws/awserr"
)

// Go get all the data of all the volumes
func GrabAllVolumesData(svc *ec2.EC2) (volume *ec2.DescribeVolumesOutput){
  volumes, err := svc.DescribeVolumes(&ec2.DescribeVolumesInput{})
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
func RemoveOldEBS(svc *ec2.EC2, input *ec2.DeleteVolumeInput) {

  result, err := svc.DeleteVolume(input)
  if err != nil {
      if aerr, ok := err.(awserr.Error); ok {
          switch aerr.Code() {
          default:
              fmt.Println(aerr.Error())
          }
      } else {
          // Print the error, cast err to awserr.Error to get the Code and
          // Message from an error.
          fmt.Println(err.Error())
      }
      return
  }

  fmt.Println(result)
}

// Go get
