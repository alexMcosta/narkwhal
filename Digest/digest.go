package digest

import (
  "fmt"

  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/aws/aws-sdk-go/aws/awserr"
)

func GrabAllVolumesData(svc *ec2.EC2) (volume *ec2.DescribeVolumesOutput){

  // Output the results of of the volumes
  volumes, err := svc.DescribeVolumes(&ec2.DescribeVolumesInput{})
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

  return volumes
}
