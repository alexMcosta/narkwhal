package main

import (
  "fmt"
  "encoding/json"
  "time"

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

  // Grab the data from AWS and Marshal it
  data, _ := json.Marshal(grabVolumes(svc))

  // Output the struct of the marshaled AWS Data
  fmt.Println(parseIt(data))

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

func parseIt(data []byte) interface{}{
  // Turn data into a string for easier ingesting
  //var dataByteArray = []byte(data.GoString())

  //Make a Volume struct for JSON parsing
  type Attachments struct {
    AttachTime          *time.Time
    DeleteOnTermination bool
    Device              string
    InstanceId          string
    State               string
    VolumeId            string
  }

  type Tags struct {
    Key   string `JSON:"Key"`
    Value string `JSON:"Value"`
  }

  type Volumes struct {
    Attachments      []Attachments
    AvailabilityZone string     `JSON:"AvailabilityZone"`
    CreateTime       *time.Time `JSON:"CreateTime"`
    Encrypted        bool       `JSON:"Encrypted"`
    Iops             int        `JSON:"Iops"`
    Size             int        `JSON:"Size"`
    SnapshotId       string     `JSON:"SnapshotId"`
    State            string     `JSON:"State"`
    Tags             []Tags
    VolumeId         string     `JSON:"VolumeId"`
    VolumeType       string     `JSON:"VolumeType"`
  }

  type Volume struct {
    Volumes []Volumes
	}

  var dataVolumes Volume

  err := json.Unmarshal(data, &dataVolumes)
  if err != nil {
    fmt.Println("error:", err)
  }

  return dataVolumes
}
