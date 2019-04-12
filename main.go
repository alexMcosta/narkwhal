package main

import (
  "fmt"

  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/session"
  )

func main() {

  // First let us configure the session
  // This session will take the default credentials found at ~/.aws/
  sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
    }))

  // Create a service with our session
  svc := ec2.New(sess)

  // Output the results of of the volumes
  fmt.Println(grabVolumes(svc))
}

func grabVolumes(svc *ec2.EC2) (volume *ec2.DescribeVolumesOutput){

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
